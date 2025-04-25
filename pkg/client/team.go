package client

import (
	"context"
	"encoding/json"
	"fmt"
	oapigen "github.com/pmurley/go-mlb/pkg/gen"
	"github.com/pmurley/go-mlb/pkg/model"
	"net/http"
)

func (c *Client) GetTeams(ctx context.Context, teamParams *oapigen.TeamsParams) ([]model.Team, error) {
	resp, err := c.apiClient.TeamsWithResponse(ctx, teamParams)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("received unexpected response status: %d: %s", resp.StatusCode(), resp.Status())
	}

	var result model.TeamsResponse
	err = json.Unmarshal(resp.Body, &result)
	if err != nil {
		return nil, err
	}

	return result.Teams, nil
}
