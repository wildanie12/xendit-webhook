package handlers

import (
	"io"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type WebhookHandler struct {
}

func NewWebhook() *WebhookHandler {
	return &WebhookHandler{}
}


func (wh *WebhookHandler) FVAPaid(c echo.Context) error {

	req := c.Request().Body

	b, err := io.ReadAll(req)
	if err != nil {
		log.Println("err", err)
	}

	log.Println("----------------------")
	log.Println(" [v] received FVA Paid request ")
	log.Println("response: ")
	log.Println(string(b))

	return c.JSON(http.StatusOK, map[string]interface{} {
		"status": "OK",
		"code": http.StatusOK,
		"message": "webhook received",
	})
}

func (wh *WebhookHandler) FVACreated(c echo.Context) error {
	req := c.Request().Body

	b, err := io.ReadAll(req)
	if err != nil {
		log.Println("err", err)
	}

	log.Println("----------------------")
	log.Println(" [v] received FVA Created request ")
	log.Println("response: ")
	log.Println(string(b))

	return c.JSON(http.StatusOK, map[string]interface{} {
		"status": "OK",
		"code": http.StatusOK,
		"message": "webhook received",
	})
}

func (wh *WebhookHandler) PaymentSucceeded(c echo.Context) error {
	req := c.Request().Body

	b, err := io.ReadAll(req)
	if err != nil {
		log.Println("err", err)
	}

	log.Println("----------------------")
	log.Println(" [v] received PaymentSucceeded request ")
	log.Println("response: ")
	log.Println(string(b))

	return c.JSON(http.StatusOK, map[string]interface{} {
		"status": "OK",
		"code": http.StatusOK,
		"message": "webhook received",
	})
}

func (wh *WebhookHandler) PaymentAwaitingCapture(c echo.Context) error {
	req := c.Request().Body

	b, err := io.ReadAll(req)
	if err != nil {
		log.Println("err", err)
	}

	log.Println("----------------------")
	log.Println(" [v] received PaymentAwaitingCapture request ")
	log.Println("response: ")
	log.Println(string(b))

	return c.JSON(http.StatusOK, map[string]interface{} {
		"status": "OK",
		"code": http.StatusOK,
		"message": "webhook received",
	})
}

func (wh *WebhookHandler) PaymentPending(c echo.Context) error {
	req := c.Request().Body

	b, err := io.ReadAll(req)
	if err != nil {
		log.Println("err", err)
	}

	log.Println("----------------------")
	log.Println(" [v] received PaymentPending request ")
	log.Println("response: ")
	log.Println(string(b))

	return c.JSON(http.StatusOK, map[string]interface{} {
		"status": "OK",
		"code": http.StatusOK,
		"message": "webhook received",
	})
}

func (wh *WebhookHandler) PaymentFailed(c echo.Context) error {
	req := c.Request().Body

	b, err := io.ReadAll(req)
	if err != nil {
		log.Println("err", err)
	}

	log.Println("----------------------")
	log.Println(" [v] received PaymentFailed request ")
	log.Println("response: ")
	log.Println(string(b))

	return c.JSON(http.StatusOK, map[string]interface{} {
		"status": "OK",
		"code": http.StatusOK,
		"message": "webhook received",
	})
}

func (wh *WebhookHandler) CapturedSucceeded(c echo.Context) error {
	req := c.Request().Body

	b, err := io.ReadAll(req)
	if err != nil {
		log.Println("err", err)
	}

	log.Println("----------------------")
	log.Println(" [v] received CapturedSucceeded request ")
	log.Println("response: ")
	log.Println(string(b))

	return c.JSON(http.StatusOK, map[string]interface{} {
		"status": "OK",
		"code": http.StatusOK,
		"message": "webhook received",
	})
}

func (wh *WebhookHandler) CapturedFailed(c echo.Context) error {
	req := c.Request().Body

	b, err := io.ReadAll(req)
	if err != nil {
		log.Println("err", err)
	}

	log.Println("----------------------")
	log.Println(" [v] received CapturedFailed request ")
	log.Println("response: ")
	log.Println(string(b))

	return c.JSON(http.StatusOK, map[string]interface{} {
		"status": "OK",
		"code": http.StatusOK,
		"message": "webhook received",
	})
}