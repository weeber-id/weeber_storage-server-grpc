package file

import (
	"errors"
	"net/url"
	"path"
	"strings"

	"github.com/weeber-id/weeber_storage-server-grpc/lib/tools"
	prbs "github.com/weeber-id/weeber_storage-server-grpc/protobuf/v1/PrivateStorage"
	pbs "github.com/weeber-id/weeber_storage-server-grpc/protobuf/v1/PublicStorage"
)

// MinioObj struct to immediatelly from protobuf and minio client
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

// FromPublicFile protobuf
func (f *MinioObj) FromPublicFile(inputFile *pbs.File) {
	f.BucketName = "public"
	f.ObjectName = path.Join(inputFile.Projectname, inputFile.Objectname)
	f.File = inputFile.File

	if inputFile.Option != nil {
		f.Option.ContentType = inputFile.Option.ContentType
	} else {
		f.Option.ContentType = tools.GetContentTypeFromExt(path.Ext(f.ObjectName))
	}
}

// FromFileLocation protobuf
// example input: projectName/objectName from protobuf
func (f *MinioObj) FromFileLocation(input *prbs.FileLocation) {
	paths := strings.Split(input.Location, "/")

	f.BucketName = paths[0]
	f.ObjectName = strings.Join(paths[1:], "/")
}

// FromURL protobuf
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

// MinioObjPrivate struct to immediatelly from protobuf and minio client
// For private file
type MinioObjPrivate struct {
	MinioObj
	Location string
}

// FromPrivateFile protobuf
func (f *MinioObjPrivate) FromPrivateFile(inputFile *prbs.File) {
	f.BucketName = inputFile.Projectname
	f.ObjectName = inputFile.Objectname
	f.Location = path.Join(inputFile.Projectname, inputFile.Objectname)
	f.File = inputFile.File

	if inputFile.Option != nil {
		f.Option.ContentType = inputFile.Option.ContentType
	} else {
		f.Option.ContentType = tools.GetContentTypeFromExt(path.Ext(f.ObjectName))
	}
}
