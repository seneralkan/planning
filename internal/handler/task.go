package handler

import (
	"planning/internal/resource/request"
	"planning/internal/resource/response"
	"planning/internal/services"

	"github.com/gofiber/fiber/v2"
)

type ITaskHandler interface {
	DistributeTasks(c *fiber.Ctx) error
}

type taskHandler struct {
	taskService services.ITaskService
}

func NewTaskHandler(taskService services.ITaskService) ITaskHandler {
	return &taskHandler{
		taskService: taskService,
	}
}

func (t *taskHandler) DistributeTasks(c *fiber.Ctx) error {
	var req request.Developers
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	tasks, totalWeeks, err := t.taskService.DistributeTasks(req.Developers)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response := response.TaskSchedule{
		Schedule:   tasks,
		TotalWeeks: totalWeeks,
	}

	return c.JSON(response)
}
