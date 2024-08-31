package app

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/arun-murali0/go-lang-backend/config"
	"github.com/arun-murali0/go-lang-backend/router"
)

func SetupAndRunApp() error {

	// load Env
	err := config.DotEnv()
	if err != nil {
		return err
	}

	// fiber
	app := fiber.New()

	// middleware
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path} ${latency}\n",
	}))

	// routes
	router.SetUpRoutes(app)

	// get the port and start
	port := os.Getenv("PORT")
	app.Listen(":" + port)

	return nil
}
