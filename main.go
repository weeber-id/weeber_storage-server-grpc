package main

import (
	"log"
	"net"

	"github.com/weeber-id/weeber_storage-server-grpc/lib/controller"
	"github.com/weeber-id/weeber_storage-server-grpc/lib/services"
	"github.com/weeber-id/weeber_storage-server-grpc/storage"
	"github.com/weeber-id/weeber_storage-server-grpc/variable"
	"google.golang.org/grpc"
)

func main() {
	variable.Initialization()
	services.MinioInitialization()

	addr := "localhost:3000"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	s := controller.Server{}
	storage.RegisterPublicStorageServer(grpcServer, &s)

	log.Printf("Starting GRPC Weeber Storage Server on %s", addr)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed GRPC listen in %s: %v", addr, err)
	}
}
