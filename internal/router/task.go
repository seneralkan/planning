package router

import "github.com/gofiber/fiber/v2"

func (r *router) RegisterTaskRoutes(router fiber.Router) {
	router.Post("tasks/schedule", r.taskHandler.DistributeTasks)
}
