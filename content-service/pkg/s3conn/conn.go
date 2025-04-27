package s3conn

import (
	"fmt"
	"net/http"
)

func Connect() error {
	return nil
}

func Ping() error {
	return nil
}

func CreateBucket(name string, url string, httpClient *http.Client) error {
	req, err := http.NewRequest("PUT", url+"/"+name, nil)
	if err != nil {
		return err
	}

	res, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusConflict {
		return fmt.Errorf("bucket with name %s is not created", name)
	}

	return nil
}
