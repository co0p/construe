package translate

import (
	"net/http"
	"time"
)

type GoogleClient struct {
	client *http.Client
	key    string
}

func NewGoogleClient(key string) GoogleClient {
	tr := &http.Transport{
		MaxIdleConns:    10,
		IdleConnTimeout: 30 * time.Second,
	}
	return GoogleClient{
		key:    key,
		client: &http.Client{Transport: tr},
	}
}

func (c GoogleClient) Translate(str string) (string, error) {
	return "", nil
}
