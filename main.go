package main

import (
	"fmt"
	"gecko-crawl/routes"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/joho/godotenv"
)

func main() {
	env := godotenv.Load()

	if env != nil {
		panic("Cannot find environment variables")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = ":4000"
	} else {
		port = ":" + port
	}

	app := fiber.New(fiber.Config{
		IdleTimeout: 5 * time.Second,
	})

	app.Use(compress.New())

	routes.SetRoutes(app)

	// Start the server and listen

	go func() {
		if err := app.Listen(port); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c // Block main thread until interrupted
	app.Shutdown()
	fmt.Println("==Shutting down server==")
}
