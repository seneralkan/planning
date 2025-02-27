package main

import (
	"errors"
	"log"
	"os"
	"os/signal"
	provider "planning"
	"planning/pkg/sqlite"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

const (
	// configs can be moved to a config file
	// and arranged in a config struct for better environment management
	// but for simplicity, I will leave them here
	dbName = "tasks"
	// the given URIs are mock URIs
	// and that's why I will use them directly
	// moreover, the mock URIs not returning data and it gives an error while fetching data
	provider1URI = "https://run.mocky.io/v3/27b47d79-f382-4dee-b4fe-a0976ceda9cd"
	provider2URI = "https://run.mocky.io/v3/7b0ff222-7a9c-4c54-9396-0df58e289143"
)

const (
	UnexpectedErrCode = "unexpected_error"
	UnexpectedErrMsg  = "Unexpected error"
)

func main() {
	// for simplicity, I will use the sqlite database
	// but in a real-world scenario, we can use a more robust database
	sqliteInstance, err := sqlite.NewSqliteInstance(dbName)
	if err != nil {
		log.Fatal("Failed to initialize sqlite instance", err)
	}

	defer sqliteInstance.Close()

	app := bootstrap(sqliteInstance)

	go func() {
		if err := app.Listen(":7878"); err != nil {
			log.Fatalf("Application Starting Error: %s", err.Error())
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c

	log.Println("application gracefully shutting down")
	if gShoutDown := app.Shutdown(); gShoutDown != nil {
		log.Fatal(gShoutDown)
	}

}

func bootstrap(db sqlite.ISqliteInstance) *fiber.App {
	app := fiber.New(fiber.Config{
		ReadTimeout: time.Minute * 1,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			message := UnexpectedErrMsg
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
				message = e.Message
			}
			return ctx.Status(code).JSON(fiber.Map{
				"error": message,
				"code":  UnexpectedErrCode,
			})
		},
	})
	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))
	// create client providers
	providers := provider.CreateClientProviders(provider1URI, provider2URI)

	// create and register the rest of the all dependencies
	r := provider.CreateRouter(db, providers)
	r.RegisterRoutes(app)

	return app
}
