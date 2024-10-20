package client

import (
	pb "breezeFS/breezeFS/proto"
	"bytes"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// Client represents the structure for the client to handle file operations
type Client struct {
	ManagerAddress string
}

// NewClient creates a new client with the given manager address
func NewClient(managerAddress string) *Client {
	return &Client{ManagerAddress: managerAddress}
}

// GetNodesForChunks requests the Manager Node for addresses of Data Nodes for chunk uploads
func (c *Client) GetNodesForChunks(fileID string, totalChunks int, fileType string) ([]*pb.ChunkNodeInfo, error) {
	conn, err := grpc.Dial(c.ManagerAddress, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Manager Node: %v", err)
	}
	defer conn.Close()

	client := pb.NewManagerServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &pb.GetNodesForChunksRequest{
		FileId:      fileID,
		TotalChunks: int32(totalChunks),
		FileType:    fileType,
	}

	resp, err := client.GetNodesForChunks(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to get nodes for chunks: %v", err)
	}

	return resp.Nodes, nil
}

// UploadChunk uploads a single chunk to the specified Data Nodes
// UploadChunk uploads a chunk of data to the specified nodes without using multipart/form-data.
func (c *Client) UploadChunk(chunk []byte, fileID, chunkID string, nodeAddresses []string) error {
	for _, nodeAddress := range nodeAddresses {
		// Construct the URL with query parameters to identify the file and chunk
		url := fmt.Sprintf("http://%s/upload?file_id=%s&chunk_id=%s", nodeAddress, fileID, chunkID)

		// Create an HTTP POST request with the raw chunk data
		req, err := http.NewRequest("POST", url, bytes.NewReader(chunk))
		if err != nil {
			return fmt.Errorf("failed to create request: %v", err)
		}

		// Set headers to indicate raw binary data
		req.Header.Set("Content-Type", "application/octet-stream")

		// Execute the request
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return fmt.Errorf("failed to upload chunk: %v", err)
		}
		defer resp.Body.Close()

		// Check if the server responded with a status OK
		if resp.StatusCode != http.StatusOK {
			responseBody, _ := io.ReadAll(resp.Body)
			return fmt.Errorf("error response from server: %s", responseBody)
		}

		log.Printf("Chunk %s uploaded to %s successfully", chunkID, nodeAddress)
	}
	return nil
}

// UploadFile handles splitting the file and uploading them to assigned nodes
func (c *Client) UploadFile(filePath, fileID string, chunkSize int) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return fmt.Errorf("failed to get file info: %v", err)
	}
	totalChunks := (int(fileInfo.Size()) + chunkSize - 1) / chunkSize

	// Extract the file type from the file path
	fileType := filepath.Ext(filePath) // Extracts the file extension (e.g., ".txt")
	if len(fileType) > 0 {
		fileType = fileType[1:] // Remove the leading dot (e.g., "txt")
	}

	// Get nodes for chunks from the Manager Node
	nodes, err := c.GetNodesForChunks(fileID, totalChunks, fileType)
	if err != nil {
		return fmt.Errorf("failed to get nodes for chunks: %v", err)
	}

	// Buffer for reading chunks
	buffer := make([]byte, chunkSize)

	for i := 0; i < totalChunks; i++ {
		n, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			return fmt.Errorf("failed to read file: %v", err)
		}
		if n == 0 {
			break
		}

		chunk := buffer[:n]
		chunkIDStr := fmt.Sprintf("%d", i)

		// Collect node addresses assigned to this chunk
		nodeAddresses := []string{}
		for _, node := range nodes {
			if node.ChunkId == int32(i) {
				nodeAddresses = append(nodeAddresses, node.NodeAddress)
			}
		}

		// Upload the chunk to all assigned nodes
		err = c.UploadChunk(chunk, fileID, chunkIDStr, nodeAddresses)
		if err != nil {
			return fmt.Errorf("failed to upload chunk: %v", err)
		}
	}

	log.Println("File upload completed successfully with duplication")
	return nil
}

// NEW CODE HERE

// GetChunkLocations requests the Manager Node for the locations of each chunk of the file
func (c *Client) GetChunkLocations(fileID string) ([]*pb.ChunkLocationInfo, string, error) {
	conn, err := grpc.Dial(c.ManagerAddress, grpc.WithInsecure())
	if err != nil {
		return nil, "", fmt.Errorf("failed to connect to Manager Node: %v", err)
	}
	defer conn.Close()

	client := pb.NewManagerServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &pb.GetChunkLocationsRequest{
		FileId: fileID,
	}

	resp, err := client.GetChunkLocations(ctx, req)
	if err != nil {
		return nil, "", fmt.Errorf("failed to get chunk locations: %v", err)
	}

	fmt.Println("Chunks:", resp.Chunks)
	fmt.Println("File Type:", resp.FileType)

	return resp.Chunks, resp.FileType, nil
}

// DownloadFile downloads the file by fetching each chunk from the available nodes
func (c *Client) DownloadFile(fileID string, chunkSize int) error {
	// Get chunk locations from the Manager Node
	chunkLocations, fileType, err := c.GetChunkLocations(fileID)
	if err != nil {
		return fmt.Errorf("failed to get chunk locations: %v", err)
	}

	outputFile := fmt.Sprintf("%s.%s", "output", fileType)

	// Create the output file
	outFile, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("failed to create output file: %v", err)
	}
	defer outFile.Close()

	// Download each chunk and write to the output file
	for _, chunkInfo := range chunkLocations {
		chunkData, err := c.downloadChunkFromAvailableNodes(fileID, chunkInfo)
		if err != nil {
			return fmt.Errorf("failed to download chunk %d: %v", chunkInfo.ChunkId, err)
		}

		// Calculate the offset based on the chunk ID and chunk size
		offset := int64(chunkInfo.ChunkId) * int64(chunkSize)

		// Seek to the correct position in the output file before writing
		_, err = outFile.Seek(offset, io.SeekStart)
		if err != nil {
			return fmt.Errorf("failed to seek to offset %d for chunk %d: %v", offset, chunkInfo.ChunkId, err)
		}

		// Write the chunk data to the output file at the correct offset
		_, err = outFile.Write(chunkData)
		if err != nil {
			return fmt.Errorf("failed to write chunk %d to output file: %v", chunkInfo.ChunkId, err)
		}
		log.Printf("Chunk %d downloaded and written successfully at offset %d", chunkInfo.ChunkId, offset)
	}

	log.Printf("File %s downloaded and reassembled successfully", outputFile)
	return nil
}

// downloadChunkFromAvailableNodes tries to download a chunk from any of the available nodes
func (c *Client) downloadChunkFromAvailableNodes(fileID string, chunkInfo *pb.ChunkLocationInfo) ([]byte, error) {
	for _, nodeAddress := range chunkInfo.Nodes {
		url := fmt.Sprintf("http://%s/download?file_id=%s&chunk_id=%d", nodeAddress, fileID, chunkInfo.ChunkId)
		resp, err := http.Get(url)
		if err != nil {
			log.Printf("Failed to download chunk %d from %s: %v", chunkInfo.ChunkId, nodeAddress, err)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {

			data, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Printf("Failed to read chunk %d from %s: %v", chunkInfo.ChunkId, nodeAddress, err)
				continue
			}
			return data, nil
		} else {
			log.Printf("Error response from %s: %s", nodeAddress, resp.Status)
		}
	}

	return nil, fmt.Errorf("all nodes failed to provide chunk %d", chunkInfo.ChunkId)
}
