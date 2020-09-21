package controller

import (
	"bytes"
	context "context"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/minio/minio-go/v7"
	"github.com/weeber-id/weeber_storage-server-grpc/lib/file"
	"github.com/weeber-id/weeber_storage-server-grpc/lib/services"
	pbs "github.com/weeber-id/weeber_storage-server-grpc/storage"
)

// Server Struct
type Server struct{}

// Upload object to Minio Server
func (s *Server) Upload(ctx context.Context, input *pbs.File) (*pbs.FileURL, error) {
	var obj file.MinioObj
	obj.FromFile(input)

	info, err := services.MinioClient.PutObject(context.Background(), "public", obj.ObjectName, bytes.NewReader(obj.File), -1, minio.PutObjectOptions{ContentType: obj.Option.ContentType})
	if err != nil {
		return nil, err
	}
	log.Printf("Upload public file %s has been success", obj.ObjectName)

	return &pbs.FileURL{
		Url: info.Location,
	}, nil
}

// Delete object from server
func (s *Server) Delete(ctx context.Context, url *pbs.FileURL) (*emptypb.Empty, error) {
	var obj file.MinioObj
	if err := obj.FromURL(url.Url); err != nil {
		return nil, err
	}

	err := services.MinioClient.RemoveObject(context.Background(), "public", obj.ObjectName, minio.RemoveObjectOptions{})
	if err != nil {
		return nil, err
	}
	log.Printf("Delete from url: %s", obj.URL)

	return &emptypb.Empty{}, nil
}
