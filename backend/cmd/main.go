package main

import (
	"cherry-blog/cmd/config"
	"cherry-blog/cmd/router"
	"log"
	"os"
	"os/signal"
	"time"

	database "cherry-blog/db/connection"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New(config.FiberConfig)

	app.Use(cors.New())
	app.Use(recover.New())
	app.Use(compress.New())
	app.Use(logger.New(logger.Config{
		TimeFormat: time.DateTime,
	}))

	// Create connection to database
	connection := database.Connection()

	// Initialize routes
	router.InitializeRouters(app, connection)

	// Start listening on port 8000
	go func() {
		if err := app.Listen(":8000"); err != nil {
			panic(err)
		}
	}()

	// Graceful shutdown
	GracefulShutdown(app, connection)
}

func GracefulShutdown(app *fiber.App, connection *gorm.DB) {

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	log.Println("Received terminate,graceful shutdown", sig)

	// Graceful shutdown for db connection
	database, err := connection.DB()
	if err != nil {
		log.Fatalln("DB Closing ERROR :", err)
	}
	database.Close()
	log.Println("DB has been closed")

	app.Shutdown()
}
