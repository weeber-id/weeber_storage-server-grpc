package services

import (
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/weeber-id/weeber_storage-server-grpc/variable"
)

// MinioClient variable
var MinioClient *minio.Client

// MinioInitialization for connect to minio server
func MinioInitialization() {
	var err error
	config := variable.MinioConfig

	MinioClient, err = minio.New(config.EndPoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AccessKey, config.SecretKey, ""),
		Secure: true,
	})
	if err != nil {
		log.Fatalf("Cannot connect to Minio Server: %v", err)
	}

}
