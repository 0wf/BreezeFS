package main

import (
	pb "breezeFS/breezeFS/proto"
	"breezeFS/internal/server"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	// Set up a listener on port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	// Register the ManagerNode service with the gRPC server
	pb.RegisterManagerServiceServer(grpcServer, server.NewManagerNode())

	log.Println("Manager Node is running on port 50051")
	// Start serving incoming connections
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
