package main

import (
	"breezeFS/internal/server"
)

func main() {
	// Define the addresses
	managerAddress := "localhost:50051" // Address of the Manager Node
	dataNodeAddress := "localhost"      // Address of this Data Node

	// Create a new Data Node instance
	dataNode := server.NewDataNode(managerAddress, dataNodeAddress)

	// Start the HTTP server for chunk operations
	dataNode.StartHTTPServer()
}
