package storage

import (
	context "context"
	"log"
)

type Server struct{}

func (s *Server) Upload(ctx context.Context, file *File) (*FilePublicInformation, error) {
	log.Printf("Receive file: %s", file.Filename)

	return &FilePublicInformation{
		Filename: file.Filename,
		Url:      "https://storages.weeber.id",
	}, nil
}
