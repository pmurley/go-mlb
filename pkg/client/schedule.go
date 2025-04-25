package client

import (
	"context"
	"encoding/json"
	"fmt"
	oapigen "github.com/pmurley/go-mlb/pkg/gen"
	"github.com/pmurley/go-mlb/pkg/model"
)

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
