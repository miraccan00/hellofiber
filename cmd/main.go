package main

import (
    "log"

    "github.com/gofiber/fiber/v2"
    "github.com/miraccan00/cookieserver/internal/handler"
)

func main() {
    app := fiber.New()
    app.Get("/hello", handler.Hello)

    if err := app.Listen(":8080"); err != nil {
        log.Fatal(err)
    }
}