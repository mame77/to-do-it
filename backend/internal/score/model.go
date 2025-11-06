package score

import "github.com/mame77/to-do-it/backend/internal/model"

type Motivation = model.Motivation

// PlayResult remains local to score package
type PlayResult struct {
	ScheduleID string `json:"schedule_id"`
	Result     string `json:"result"`
}
