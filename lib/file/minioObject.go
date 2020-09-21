package file

import (
	"errors"
	"net/url"
	"path"
	"strings"

	"github.com/weeber-id/weeber_storage-server-grpc/storage"
)

type MinioObj struct {
	URL        string
	HostName   string
	BucketName string
	ObjectName string
	File       []byte
	Option     struct {
		ContentType string
	}
}

func (f *MinioObj) FromFile(inputFile *storage.File) {
	f.ObjectName = path.Join(inputFile.Projectname, inputFile.Objectname)
	f.File = inputFile.File

	if inputFile.Option != nil {
		f.Option.ContentType = inputFile.Option.ContentType
	}
}

func (f *MinioObj) FromURL(URI string) error {
	var err error

	if len(URI) < 8 {
		return errors.New("Invalid length URL")
	}

	// if url[:7] != "https://" {
	// 	return errors.New("Invalid URL")
	// } else if url[:6] != "http://" {
	// 	return errors.New("Invalid URL")
	// }

	URIs := strings.Split(URI, "/")
	if len(URIs) < 4 {
		return errors.New("Invalid length path URL")
	}

	f.URL = URI
	f.HostName = URIs[3]
	f.BucketName = URIs[4]
	f.ObjectName, err = url.QueryUnescape(strings.Join(URIs[4:], "/"))
	if err != nil {
		return err
	}
	return nil
}
