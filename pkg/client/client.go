package client

import (
	oapigen "github.com/pmurley/go-mlb/pkg/gen"
)

type Client struct {
	apiClient *oapigen.ClientWithResponses
}

func NewClient(baseURL string, options ...oapigen.ClientOption) (*Client, error) {
	client, err := oapigen.NewClientWithResponses(baseURL, options...)
	if err != nil {
		return nil, err
	}
	return &Client{
		apiClient: client,
	}, nil
}
