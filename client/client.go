package main

import (
	"context"
	"log"

	"github.com/weeber-id/weeber_storage-server-grpc/storage"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Cannot connect server: %v", err)
	}
	defer conn.Close()

	client := storage.NewPublicStorageClient(conn)

	response, err := client.Upload(context.Background(), &storage.File{Filename: "weeber.jpg"})
	if err != nil {
		log.Fatalf("Cannot upload file: %v", err)
	}
	log.Printf("Response from server: %v", response)
}
