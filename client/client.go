package main

import (
	"bytes"
	"context"
	"log"
	"os"

	prbs "github.com/weeber-id/weeber_storage-server-grpc/protobuf/v1/PrivateStorage"
	pbs "github.com/weeber-id/weeber_storage-server-grpc/protobuf/v1/PublicStorage"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:3000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Cannot connect server: %v", err)
	}
	defer conn.Close()

	client := pbs.NewPublicStorageClient(conn)
	clientPrivate := prbs.NewPrivateStorageClient(conn)

	// Read local file
	data, err := os.Open("../example/unnamed.jpg")
	if err != nil {
		log.Fatalf("Error when get local image %v", err)
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(data)

	// Upload public file
	fileInput := &pbs.File{
		File:        buf.Bytes(),
		Projectname: "vokasi-binadesa",
		Objectname:  "testing.jpg",
	}
	resp, err := client.Upload(context.Background(), fileInput)
	if err != nil {
		log.Fatalf("Error from server %v", err)
	}
	log.Printf("Link public file %s", resp.Url)

	// Delete public file from URL
	file := &pbs.FileURL{
		Url: "https://storages.weeber.id/public/vokasi_binadesa/testing.jpg",
	}
	_, err = client.Delete(context.Background(), file)
	if err != nil {
		log.Fatalf("Cannot remove file: %v", err)
	}

	// Upload private file
	inputUpload := &prbs.File{
		File:        buf.Bytes(),
		Projectname: "vokasi-binadesa",
		Objectname:  "contoh/testing.jpg",
	}
	loc, err := clientPrivate.Upload(context.Background(), inputUpload)
	if err != nil {
		log.Fatalf("error in upload private file %v", err)
	}
	log.Printf("Location private file: %s", loc.Location)

	// Download private file
	inputDownload := &prbs.FileLocation{
		Location: loc.Location,
	}
	dataPrivate, err := clientPrivate.Download(context.Background(), inputDownload)
	if err != nil {
		log.Fatalf("Error in download private file %v", dataPrivate)
	}
	log.Printf("Size download private file %d bytes", len(dataPrivate.File))

	// Delete private file
	inputDelete := &prbs.FileLocation{
		Location: loc.Location,
	}
	_, err = clientPrivate.Delete(context.Background(), inputDelete)
	if err != nil {
		log.Printf("Error in delete private file %v", err)
	}
}
