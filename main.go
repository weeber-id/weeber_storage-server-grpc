package main

import (
	"log"
	"net"

	"github.com/weeber-id/weeber_storage-server-grpc/storage"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	s := storage.Server{}
	storage.RegisterPublicStorageServer(grpcServer, &s)

	log.Println("Starting GRPC Weeber Storage Server")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed GRPC listen in port 9000: %v", err)
	}
}
