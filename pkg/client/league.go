package client

import (
	"context"
	"encoding/json"
	"fmt"
	oapigen "github.com/pmurley/go-mlb/pkg/gen"
	"github.com/pmurley/go-mlb/pkg/model"
)

type LeagueId int32

// GetLeagues allows for fetching data about leagues.
// I haven't found a way to get a LeagueID for all of MLB, but AL is 103 and NL is 104
func (c *Client) GetLeagues(ctx context.Context, params *oapigen.LeagueParams) ([]model.League, error) {
	resp, err := c.apiClient.LeagueWithResponse(ctx, params)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	var leagues model.LeaguesResponse
	err = json.Unmarshal(resp.Body, &leagues)
	if err != nil {
		return nil, err
	}

	return leagues.Leagues, nil
}

// GetAllStarBallot allows for fetching data about the All Star ballot.
// It appears these all star endpoint don't actually return vote data currently, just player info
func (c *Client) GetAllStarBallot(ctx context.Context, params *oapigen.AllStarBallotParams) ([]model.PotentialAllStarPlayer, error) {
	resp, err := c.apiClient.AllStarBallotWithResponse(ctx, params)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	var writeIns model.AllStarBallotResponse
	err = json.Unmarshal(resp.Body, &writeIns)
	if err != nil {
		return nil, err
	}

	return writeIns.People, nil
}

// GetAllStarWriteIns allows for fetching data about the All Star write ins.
// It appears these all star endpoint don't actually return vote data currently, just player info
func (c *Client) GetAllStarWriteIns(ctx context.Context, leagueId int, params *oapigen.AllStarWriteInsParams) ([]model.PotentialAllStarPlayer, error) {
	resp, err := c.apiClient.AllStarWriteInsWithResponse(ctx, int32(leagueId), params)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	var writeIns model.AllStarWriteInsResponse
	err = json.Unmarshal(resp.Body, &writeIns)
	if err != nil {
		return nil, err
	}

	return writeIns.People, nil
}

// GetAllStarFinalVote allows for fetching data about the All Star final vote.
// It appears these all star endpoint don't actually return vote data currently, just player info
func (c *Client) GetAllStarFinalVote(ctx context.Context, leagueId int, params *oapigen.AllStarFinalVoteParams) ([]model.PotentialAllStarPlayer, error) {
	resp, err := c.apiClient.AllStarFinalVoteWithResponse(ctx, int32(leagueId), params)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	var writeIns model.AllStarWriteInsResponse
	err = json.Unmarshal(resp.Body, &writeIns)
	if err != nil {
		return nil, err
	}

	return writeIns.People, nil
}
