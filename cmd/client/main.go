package main

import (
	"breezeFS/internal/client"
	"flag"
	"log"
	"path/filepath"
	"strings"
)

func main() {

	operation := flag.String("op", "upload", "Operation type: upload or download")
	filePath := flag.String("filepath", "", "Path to the file to upload or download")
	chunkSize := flag.Int("chunksize", 1024*1024, "Size of each chunk in bytes (default 1MB)")

	// Parse the flags
	flag.Parse()

	// Initialize the client with the Manager Node address
	client := client.NewClient("localhost:50051")

	switch *operation {
	case "upload":
		// Upload the file
		fileNameWithExt := filepath.Base(*filePath)
		fileID := strings.TrimSuffix(fileNameWithExt, filepath.Ext(fileNameWithExt))
		if err := client.UploadFile(*filePath, fileID, *chunkSize); err != nil { // Example chunk size: 1MB
			log.Fatalf("Failed to upload file: %v", err)
		}
		log.Println("File uploaded successfully")

	case "download":
		// Validate output path for download
		fileNameWithExt := filepath.Base(*filePath)
		fileID := strings.TrimSuffix(fileNameWithExt, filepath.Ext(fileNameWithExt))

		// Download the file
		if err := client.DownloadFile(fileID, *chunkSize); err != nil {
			log.Fatalf("Failed to download file: %v", err)
		}
		log.Println("File downloaded successfully")

	default:
		log.Fatalf("Invalid operation: %s. Use 'upload' or 'download'", *operation)
	}
}
