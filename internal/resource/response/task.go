package response

import "planning/internal/models"

type TaskSchedule struct {
	Schedule   map[string][]models.Task `json:"schedule"`
	TotalWeeks int                      `json:"totalWeeks"`
}
