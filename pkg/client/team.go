package client

import (
	"context"
	"encoding/json"
	"fmt"
	oapigen "github.com/pmurley/go-mlb/pkg/gen"
	"net/http"
)

func (c *Client) GetTeams(ctx context.Context, teamParams *oapigen.TeamsParams) ([]Team, error) {
	resp, err := c.apiClient.TeamsWithResponse(ctx, teamParams)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("received unexpected response status: %d: %s", resp.StatusCode(), resp.Status())
	}

	var result TeamsResponse
	err = json.Unmarshal(resp.Body, &result)
	if err != nil {
		return nil, err
	}

	return result.Teams, nil
}
