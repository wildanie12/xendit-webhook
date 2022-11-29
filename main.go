package main

import (
	"os"
	"xendit-sandbox/handlers"
	"xendit-sandbox/routers"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	wh := handlers.NewWebhook()

	e.Pre(middleware.RemoveTrailingSlash())
	routers.SetPaymentRoutes(e, wh)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8081"
	}
	e.Logger.Fatal(e.Start(":" + port))
}