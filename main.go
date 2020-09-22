package main

import (
	"log"
	"net"

	"github.com/weeber-id/weeber_storage-server-grpc/lib/controller"
	"github.com/weeber-id/weeber_storage-server-grpc/lib/services"
	prbs "github.com/weeber-id/weeber_storage-server-grpc/protobuf/v1/PrivateStorage"
	pbs "github.com/weeber-id/weeber_storage-server-grpc/protobuf/v1/PublicStorage"
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

	publicStorageCtl := controller.PublicStorageServer{}
	pbs.RegisterPublicStorageServer(grpcServer, &publicStorageCtl)

	privateStorageCtl := controller.PrivateStorageServer{}
	prbs.RegisterPrivateStorageServer(grpcServer, &privateStorageCtl)

	log.Printf("Starting GRPC Weeber Storage Server on %s", addr)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed GRPC listen in %s: %v", addr, err)
	}
}
