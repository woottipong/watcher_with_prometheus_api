package networking

import (
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

var client = &http.Client{Timeout: 15 * time.Second}

// GetContent - get JSON content from API and response via given interface
func GetContent(uri string) ([]byte, error) {
	if uri == "" {
		return nil, errors.New("URI cannot be nil")
	}

	var data []byte
	response, err := client.Get(uri)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	data, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}
