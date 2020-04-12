package services

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Style struct {
	Code    string
	Message string
	Data    interface{}
}

func EchoJson(c echo.Context, s Style) error {
	return c.JSON(http.StatusOK, s)
}
