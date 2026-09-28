package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/miraccan00/cookieserver/internal/handler"
)

// version is set at build time: -ldflags "-X main.version=sha-abc1234"
var version = "dev"

func main() {
	h := handler.New(os.Getenv("MESSAGE"), os.Getenv("APP_ENV"), version)

	app := fiber.New(fiber.Config{DisableStartupMessage: true})
	app.Get("/hello", h.Hello)
	app.Get("/healthz", handler.Healthz)

	log.Printf("hellofiber %s listening on :8080", version)
	if err := app.Listen(":8080"); err != nil {
		log.Fatal(err)
	}
}
