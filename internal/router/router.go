package router

import (
	"planning/internal/handler"

	"github.com/gofiber/fiber/v2"
)

type IRouter interface {
	RegisterRoutes(app *fiber.App)
}

type router struct {
	taskHandler handler.ITaskHandler
}

func NewRouter(taskHandler handler.ITaskHandler) IRouter {
	return &router{
		taskHandler: taskHandler,
	}
}

func (r *router) RegisterRoutes(app *fiber.App) {
	// Serve static files first
	app.Static("/", "./web/static")
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("./web/templates/index.html")
	})

	// API routes
	api := app.Group("/api")
	r.RegisterTaskRoutes(api)
}
