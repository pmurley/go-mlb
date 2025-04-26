package client

import (
	"context"
	"encoding/json"
	"fmt"
	oapigen "github.com/pmurley/go-mlb/pkg/gen"
	"github.com/pmurley/go-mlb/pkg/model"
)

/***
 * Notes:
 * Tracking events endpoint is authenticated and not available
 */

func (c *Client) GetSchedule(ctx context.Context, params *oapigen.ScheduleParams) (*model.ScheduleResponse, error) {
	resp, err := c.apiClient.ScheduleWithResponse(ctx, params)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("unexpected status code: %d %s", resp.StatusCode(), resp.Status())
	}

	var schedule model.ScheduleResponse
	err = json.Unmarshal(resp.Body, &schedule)
	if err != nil {
		return nil, err
	}

	return &schedule, nil
}

func (c *Client) GetPostseasonsSchedule(ctx context.Context, params *oapigen.PostseasonScheduleParams) (*model.PostseasonScheduleResponse, error) {
	resp, err := c.apiClient.PostseasonScheduleWithResponse(ctx, params)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	var postseasonSchedule model.PostseasonScheduleResponse
	err = json.Unmarshal(resp.Body, &postseasonSchedule)
	if err != nil {
		return nil, err
	}

	return &postseasonSchedule, nil
}

func (c *Client) GetPostseasonTuneInSchedule(ctx context.Context, params *oapigen.TuneInParams) (*model.TuneInScheduleResponse, error) {
	resp, err := c.apiClient.TuneInWithResponse(ctx, params)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	var postseasonTuneInSchedule model.TuneInScheduleResponse
	err = json.Unmarshal(resp.Body, &postseasonTuneInSchedule)
	if err != nil {
		return nil, err
	}

	return &postseasonTuneInSchedule, nil
}

func (c *Client) GetPostseasonSeriesSchedule(ctx context.Context, params *oapigen.PostseasonScheduleSeriesParams) (*model.PostseasonSeriesScheduleResponse, error) {
	resp, err := c.apiClient.PostseasonScheduleSeriesWithResponse(ctx, params)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	var postseasonSeries model.PostseasonSeriesScheduleResponse
	err = json.Unmarshal(resp.Body, &postseasonSeries)
	if err != nil {
		return nil, err
	}

	return &postseasonSeries, nil
}

func (c *Client) GetTieGamesSchedule(ctx context.Context, params *oapigen.TieGamesParams) (*model.TieGamesResponse, error) {
	resp, err := c.apiClient.TieGamesWithResponse(ctx, params)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	var tiedGames model.TieGamesResponse
	err = json.Unmarshal(resp.Body, &tiedGames)
	if err != nil {
		return nil, err
	}

	return &tiedGames, nil
}
