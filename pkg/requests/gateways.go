package requests

import (
	"net/http"
	"time"
)

type ApiConfig struct {
	BaseURL string
	Token   string
	Client  *http.Client
}

func NewGateway(baseURL, token string) *ApiConfig {
	return &ApiConfig{
		BaseURL: baseURL,
		Token:   token,
		Client: &http.Client{
			Timeout: 10 * time.Second,
			Transport: &CustomTransport{
				Token: token,
				Base:  http.DefaultTransport,
			},
		},
	}
}
