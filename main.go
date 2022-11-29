package main

import (
	"xendit-sandbox/handlers"
	"xendit-sandbox/routers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	wh := handlers.NewWebhook()

	e.Pre(middleware.RemoveTrailingSlash())
	routers.SetPaymentRoutes(e, wh)

	e.Logger.Fatal(e.Start(":8081"))
}