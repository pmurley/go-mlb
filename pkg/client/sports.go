package client

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pmurley/go-mlb/pkg/gen"
	"github.com/pmurley/go-mlb/pkg/model"
)

func (c *Client) GetSports(ctx context.Context, params *gen.SportsParams) (*[]model.Sport, error) {
	resp, err := c.apiClient.SportsWithResponse(ctx, params)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("error: unexpected response status %s", resp.Status())
	}

	var sports model.SportsResponse
	err = json.Unmarshal(resp.Body, &sports)
	if err != nil {
		return nil, err
	}

	return &sports.Sports, nil
}
func (c *Client) GetSportsPlayers(ctx context.Context, sportId int32, params *gen.SportPlayersParams) ([]model.SportsPlayer, error) {
	resp, err := c.apiClient.SportPlayersWithResponse(ctx, sportId, params)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("error: unexpected response status %s", resp.Status())
	}

	var sports model.SportsPlayerResponse
	err = json.Unmarshal(resp.Body, &sports)
	if err != nil {
		return nil, err
	}

	return sports.People, nil
}
func (c *Client) GetSportsAllSportBallot(ctx context.Context, sportId int32, params *gen.AllSportBallotParams) ([]model.SportsPlayer, error) {
	resp, err := c.apiClient.AllSportBallotWithResponse(ctx, sportId, params)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("error: unexpected response status %s", resp.Status())
	}

	var sports model.SportsPlayerResponse
	err = json.Unmarshal(resp.Body, &sports)
	if err != nil {
		return nil, err
	}

	return sports.People, nil
}
