package controller

import (
	"bytes"
	context "context"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/minio/minio-go/v7"
	"github.com/weeber-id/weeber_storage-server-grpc/lib/file"
	"github.com/weeber-id/weeber_storage-server-grpc/lib/services"
	prbs "github.com/weeber-id/weeber_storage-server-grpc/protobuf/v1/PrivateStorage"
)

// PrivateStorageServer Struct
type PrivateStorageServer struct{}

// Upload object to Minio Server
func (s *PrivateStorageServer) Upload(ctx context.Context, input *prbs.File) (*prbs.FileLocation, error) {
	var obj file.MinioObjPrivate
	obj.FromPrivateFile(input)

	_, err := services.MinioClient.PutObject(context.Background(), obj.BucketName, obj.ObjectName, bytes.NewReader(obj.File), -1, minio.PutObjectOptions{ContentType: obj.Option.ContentType})
	if err != nil {
		return nil, err
	}
	log.Printf("Upload private file %s has been success", obj.ObjectName)

	return &prbs.FileLocation{
		Location: obj.Location,
	}, nil
}

// Download private object from Minio Server
func (s *PrivateStorageServer) Download(ctx context.Context, input *prbs.FileLocation) (*prbs.File, error) {
	var obj file.MinioObjPrivate
	obj.FromFileLocation(input)

	reader, err := services.MinioClient.GetObject(context.Background(), obj.BucketName, obj.ObjectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)
	log.Printf("Download private data %v \n", input.Location)

	return &prbs.File{
		Projectname: obj.BucketName,
		Objectname:  obj.ObjectName,
		File:        buf.Bytes(),
	}, nil
}

// Delete object from server
func (s *PrivateStorageServer) Delete(ctx context.Context, input *prbs.FileLocation) (*emptypb.Empty, error) {
	var obj file.MinioObjPrivate
	obj.FromFileLocation(input)

	err := services.MinioClient.RemoveObject(context.Background(), obj.BucketName, obj.ObjectName, minio.RemoveObjectOptions{})
	if err != nil {
		return nil, err
	}
	log.Printf("Delete from location: %s", input.Location)

	return &emptypb.Empty{}, nil
}
