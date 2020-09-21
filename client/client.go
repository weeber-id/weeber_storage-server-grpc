package main

import (
	"bytes"
	"context"
	"log"
	"os"

	"github.com/weeber-id/weeber_storage-server-grpc/storage"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:3000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Cannot connect server: %v", err)
	}
	defer conn.Close()

	client := storage.NewPublicStorageClient(conn)

	// Upload public file
	data, err := os.Open("../example/unnamed.jpg")
	if err != nil {
		log.Fatalf("Error when get local image %v", err)
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(data)
	fileInput := &storage.File{
		File:        buf.Bytes(),
		Projectname: "vokasi_binadesa",
		Objectname:  "testing.jpg",
		Option: &storage.File_Option{
			ContentType: "image/jpg",
		},
	}
	resp, err := client.Upload(context.Background(), fileInput)
	if err != nil {
		log.Fatalf("Error from server %v", err)
	}
	log.Println(resp.Url)

	// Download public file from URL
	file := &storage.FileURL{
		Url: "https://storages.weeber.id/public/vokasi_binadesa/testing.jpg",
	}
	_, err = client.Delete(context.Background(), file)
	if err != nil {
		log.Fatalf("Cannot remove file: %v", err)
	}
}
