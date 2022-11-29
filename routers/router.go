package routers

import (
	"net/http"
	"xendit-sandbox/handlers"

	"github.com/labstack/echo/v4"
)

func SetPaymentRoutes(e *echo.Echo, h *handlers.WebhookHandler) {
	g := e.Group("/webhook/payment")
	g.POST("", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status": "OK",
			"code": http.StatusOK,
			"message": "welcome!, this is homepage of the webhook API",
		})
	})
	g.POST("/payment_succeeded", h.PaymentSucceeded)
	g.POST("/payment_awaiting_capture", h.PaymentAwaitingCapture)
	g.POST("/payment_pending", h.PaymentPending)
	g.POST("/payment_failed", h.PaymentFailed)
	g.POST("/capture_succeeded", h.CapturedSucceeded)
	g.POST("/capture_failed", h.CapturedFailed)
}