package s3

import (
	"context"
	"fmt"
	"net/http"

	"content-service/internal/adapter/repo/s3/dao"
	"content-service/internal/model"
	"content-service/pkg/s3conn"
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

func (f *File) Create(ctx context.Context, file model.File) (string, error) {
	object := dao.FromFile(file)

	url := f.url + "/" + f.bucket + "/" + object.ObjectKey
	req, err := http.NewRequest("PUT", url, object.Body)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", object.Type)

	res, err := f.client.Do(req)
	if err != nil {
		return "", err
	}

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("file is not uploaded: %d", res.StatusCode)
	}

	return object.ObjectKey, nil
}

func (f *File) Get(ctx context.Context, key string) (*model.File, error) {
	url := f.url + "/" + f.bucket + "/" + key
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := f.client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("cannot get the file: %s", res.Body)
	}

	file := dao.ToFile(res.Body, res.Header.Get("Content-Type"))

	return &file, nil
}

func (f *File) Delete(ctx context.Context, key string) error {
	url := f.url + "/" + f.bucket + "/" + key
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	res, err := f.client.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("cannot delete the file: %s", res.Body)
	}

	return nil
}
