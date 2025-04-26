package client

import (
	"context"
	"encoding/json"
	"fmt"
	oapigen "github.com/pmurley/go-mlb/pkg/gen"
	"github.com/pmurley/go-mlb/pkg/model"
)

func (c *Client) GetAttendance(ctx context.Context, params *oapigen.GetTeamAttendanceParams) (*model.AttendanceResponse, error) {
	resp, err := c.apiClient.GetTeamAttendanceWithResponse(ctx, params)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("unexpected status code: %d %s", resp.StatusCode(), resp.Status())
	}

	var attendance model.AttendanceResponse
	err = json.Unmarshal(resp.Body, &attendance)
	if err != nil {
		return nil, err
	}

	return &attendance, nil
}
