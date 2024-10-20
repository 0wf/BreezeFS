package server

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	pb "breezeFS/breezeFS/proto"
	"google.golang.org/grpc"
)

type DataNode struct {
	ManagerAddress string
	NodeAddress    string
}

var fileTypeMap = struct {
	sync.RWMutex
	data map[string]string
}{data: make(map[string]string)}

// NewDataNode creates a new instance of DataNode with specified addresses
func NewDataNode(managerAddress, nodeAddress string) *DataNode {
	return &DataNode{
		ManagerAddress: managerAddress,
		NodeAddress:    nodeAddress,
	}
}

// RegisterWithManager registers the Data Node with the Manager Node via gRPC
func (dn *DataNode) RegisterWithManager(nodeAddress string) error {
	conn, err := grpc.Dial(dn.ManagerAddress, grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("failed to connect to Manager Node: %v", err)
	}
	defer conn.Close()

	client := pb.NewManagerServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &pb.RegisterNodeRequest{
		NodeAddress: nodeAddress,
	}

	_, err = client.RegisterNode(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to register node: %v", err)
	}

	log.Printf("Data Node registered with Manager at %s", dn.ManagerAddress)
	return nil
}

// StartHTTPServer starts the HTTP server to handle chunk upload and download
func (dn *DataNode) StartHTTPServer() {
	// Listen on any available port
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:0", dn.NodeAddress))
	if err != nil {
		log.Fatalf("Failed to start listener: %v", err)
	}

	// Get the actual address including the dynamically assigned port
	nodeAddress := listener.Addr().String()
	log.Printf("Data Node is running on %s", nodeAddress)

	// Register the Data Node with the Manager Node
	if err := dn.RegisterWithManager(nodeAddress); err != nil {
		log.Fatalf("Failed to register with Manager Node: %v", err)
	}

	// Handle HTTP requests
	http.HandleFunc("/upload", dn.uploadChunkHandler)
	http.HandleFunc("/download", dn.downloadChunkHandler)
	if err := http.Serve(listener, nil); err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}

}

func (dn *DataNode) uploadChunkHandler(w http.ResponseWriter, r *http.Request) {
	// Get file and chunk IDs from query parameters
	fileID := r.URL.Query().Get("file_id")
	chunkID := r.URL.Query().Get("chunk_id")

	if fileID == "" || chunkID == "" {
		http.Error(w, "Missing file_id or chunk_id", http.StatusBadRequest)
		return
	}

	// Extract the file type from headers
	fileType := r.Header.Get("File-Type")

	// Store file type in the map
	fileTypeMap.Lock()
	if fileType != "" {
		fileTypeMap.data[fileID] = fileType
	}
	fileTypeMap.Unlock()

	// Create a file path to store the chunk
	filePath := fmt.Sprintf("data/%s_%s.chunk", fileID, chunkID)

	fmt.Println("filePath: " + filePath)

	// Create or open the file for writing
	out, err := os.Create(filePath)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create file: %v", err), http.StatusInternalServerError)
		return
	}
	defer out.Close()

	// Write the uploaded data to the file
	_, err = io.Copy(out, r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to save chunk: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Chunk %s of file %s stored successfully", chunkID, fileID)
}

// downloadChunkHandler handles downloading of chunks from the Data Node
func (dn *DataNode) downloadChunkHandler(w http.ResponseWriter, r *http.Request) {
	// Get file and chunk IDs from query parameters
	fileID := r.URL.Query().Get("file_id")
	chunkID := r.URL.Query().Get("chunk_id")

	if fileID == "" || chunkID == "" {
		http.Error(w, "Missing file_id or chunk_id", http.StatusBadRequest)
		return
	}

	// Construct the file path based on file and chunk IDs
	filePath := filepath.Join("data", fmt.Sprintf("%s_%s.chunk", fileID, chunkID))

	// Open the chunk file
	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to open chunk file: %v", err), http.StatusNotFound)
		return
	}
	defer file.Close()

	// Set the headers to indicate a file download
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s_%s.chunk\"", fileID, chunkID))

	// Copy the file data to the response
	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to send chunk data: %v", err), http.StatusInternalServerError)
		return
	}
}

func getFileType(fileID string) string {
	fileTypeMap.RLock()
	defer fileTypeMap.RUnlock()
	return fileTypeMap.data[fileID]
}
