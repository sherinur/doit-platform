package s3

import (
	"content-service/pkg/s3conn"
	"context"
	"fmt"
	"net/http"
)

type File struct {
	url    string
	client *http.Client
	bucket string
}

const bucketName = "files"

func NewFile(urlBase string) (*File, error) {
	httpClient := &http.Client{}
	err := s3conn.CreateBucket(bucketName, urlBase, httpClient)
	if err != nil {
		return nil, fmt.Errorf("repository error: %s", err.Error())
	}

	return &File{
		url:    urlBase,
		client: httpClient,
		bucket: bucketName,
	}, nil
}

func (f *File) Create(ctx context.Context, object []byte) (string, error) {
	// url := f.url + "/" + f.bucket + "/"

	// req, err := http.NewRequest("PUT", url)

	return "", nil
}

func (f *File) Get(key string) ([]byte, error) {
	return nil, nil
}

func (f *File) Delete(key string) error {
	return nil
}
