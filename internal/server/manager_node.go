package server

import (
	pb "breezeFS/breezeFS/proto"
	"context"
	"fmt"
	"sync"
)

type ManagerNode struct {
	pb.UnimplementedManagerServiceServer
	mu            sync.Mutex
	nodes         map[string]bool               // Registered nodes
	nodeAddresses []string                      // List of node addresses
	chunkMapping  map[string]map[int32][]string // FileID -> ChunkID -> []NodeAddresses
	fileTypes     map[string]string             // Map to store file types by file ID

	//chunks map[string][]pb.ChunkInfo
}

func NewManagerNode() *ManagerNode {
	return &ManagerNode{
		nodes:        make(map[string]bool),
		chunkMapping: make(map[string]map[int32][]string),
		fileTypes:    make(map[string]string),

		//chunks: make(map[string][]pb.ChunkInfo),
	}
}

func (m *ManagerNode) RegisterNode(ctx context.Context, req *pb.RegisterNodeRequest) (*pb.RegisterNodeResponse, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.nodes[req.NodeAddress] = true
	return &pb.RegisterNodeResponse{Message: "Node registered successfully"}, nil
}

func (m *ManagerNode) GetNodesForChunks(ctx context.Context, req *pb.GetNodesForChunksRequest) (*pb.GetNodesForChunksResponse, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Extract node addresses from the map keys
	nodeAddresses := make([]string, 0, len(m.nodes))
	for address := range m.nodes {
		nodeAddresses = append(nodeAddresses, address)
	}

	// Check if there are any registered nodes
	if len(nodeAddresses) == 0 {
		return nil, fmt.Errorf("no registered nodes available")
	}
	var chunkNodes []*pb.ChunkNodeInfo

	// Initialize chunk mapping for the file if not already present
	if _, exists := m.chunkMapping[req.FileId]; !exists {
		m.chunkMapping[req.FileId] = make(map[int32][]string)
	}

	// Assign each chunk to two nodes using round robin
	for i := 0; i < int(req.TotalChunks); i++ {

		firstNode := nodeAddresses[i%len(nodeAddresses)]
		var secondNode string

		// Only assign a second node if more than one node is available
		if len(nodeAddresses) > 1 {
			secondNode = nodeAddresses[(i+1)%len(nodeAddresses)]
		} else {
			secondNode = firstNode // Fall back to the same node if only one is available
		}

		// Update chunk mapping with the assigned nodes, ensuring no duplicates
		if firstNode != secondNode {
			m.chunkMapping[req.FileId][int32(i)] = []string{firstNode, secondNode}
		} else {
			m.chunkMapping[req.FileId][int32(i)] = []string{firstNode}
		}

		// Assign file type to map
		m.fileTypes[req.FileId] = req.FileType

		// Ensure chunk duplication by assigning the chunk to two nodes
		chunkNodes = append(chunkNodes, &pb.ChunkNodeInfo{
			ChunkId:     int32(i),
			NodeAddress: firstNode,
		}, &pb.ChunkNodeInfo{
			ChunkId:     int32(i),
			NodeAddress: secondNode,
		})
	}

	return &pb.GetNodesForChunksResponse{Nodes: chunkNodes}, nil
}

// GetChunkLocations provides the locations of each chunk for a file
func (m *ManagerNode) GetChunkLocations(ctx context.Context, req *pb.GetChunkLocationsRequest) (*pb.GetChunkLocationsResponse, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	chunkLocations := m.chunkMapping[req.FileId]
	if chunkLocations == nil {
		fmt.Println("COULD NOT FIND")
		return nil, fmt.Errorf("file not found")
	}

	fmt.Println("Chunk mapping for file %s: %v", req.FileId, chunkLocations)

	var chunkInfos []*pb.ChunkLocationInfo
	for chunkID, nodes := range chunkLocations {
		chunkInfos = append(chunkInfos, &pb.ChunkLocationInfo{
			ChunkId: chunkID,
			Nodes:   nodes,
		})
	}

	return &pb.GetChunkLocationsResponse{
		Chunks:   chunkInfos,
		FileType: m.fileTypes[req.FileId],
	}, nil
}
