package client

import (
	"context"
	"encoding/json"
	"fmt"
	oapigen "github.com/pmurley/go-mlb/pkg/gen"
)

func (c *Client) GetLeagues(ctx context.Context, params *oapigen.LeagueParams) ([]League, error) {
	resp, err := c.apiClient.LeagueWithResponse(ctx, params)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	var leagues LeaguesResponse
	err = json.Unmarshal(resp.Body, &leagues)
	if err != nil {
		return nil, err
	}

	return leagues.Leagues, nil
}
