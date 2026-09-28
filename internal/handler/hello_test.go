package handler

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func get(t *testing.T, h *Handler) string {
	t.Helper()
	app := fiber.New()
	app.Get("/hello", h.Hello)
	resp, err := app.Test(httptest.NewRequest("GET", "/hello", nil))
	if err != nil {
		t.Fatal(err)
	}
	b, _ := io.ReadAll(resp.Body)
	return string(b)
}

func TestHelloDefault(t *testing.T) {
	if got := get(t, New("", "", "dev")); got != "Hello, World!" {
		t.Fatalf("got %q", got)
	}
}

func TestHelloWithEnv(t *testing.T) {
	want := "Hello from helloapi [PROD] (env=production, version=sha-abc1234)"
	if got := get(t, New("Hello from helloapi [PROD]", "production", "sha-abc1234")); got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}
