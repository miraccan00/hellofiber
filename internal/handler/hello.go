package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	message string
	env     string
	version string
}

// New builds the handler from what the Deployment passes in (MESSAGE, APP_ENV)
// and the image's build-time version. Empty message falls back to the original greeting.
func New(message, env, version string) *Handler {
	if message == "" {
		message = "Hello, World!"
	}
	return &Handler{message: message, env: env, version: version}
}

func (h *Handler) Hello(c *fiber.Ctx) error {
	if h.env == "" {
		return c.SendString(h.message)
	}
	return c.SendString(fmt.Sprintf("%s (env=%s, version=%s)", h.message, h.env, h.version))
}

func Healthz(c *fiber.Ctx) error {
	return c.SendString("ok")
}
