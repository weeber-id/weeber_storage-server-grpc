package main

import (
	"log"
	"net"

	"github.com/weeber-id/weeber_storage-server-grpc/chat"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen, %v", err)
	}

	s := chat.Server{}

	grpcServer := grpc.NewServer()
	chat.RegisterChatServiveServer(grpcServer, &s)

	log.Println("Starting GRPC server")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed grpc listen in port 9000, %v", err)
	}
}
