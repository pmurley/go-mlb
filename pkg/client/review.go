package client

import (
	"context"
	"encoding/json"
	"fmt"
	oapigen "github.com/pmurley/go-mlb/pkg/gen"
)

func (c *Client) GetReview(ctx context.Context, params *oapigen.GetReviewInfoParams) (interface{}, error) {
	resp, err := c.apiClient.GetReviewInfoWithResponse(ctx, params)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	fmt.Println(string(resp.Body))

	var reviewResponse string
	err = json.Unmarshal(resp.Body, &reviewResponse)
	if err != nil {
		return nil, err
	}

	return reviewResponse, nil
}
