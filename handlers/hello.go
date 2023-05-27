package handlers

import (
	"github.com/labstack/echo/v4"
	"http-go-sandbox/types"
	"net/http"
)

func Hello(c echo.Context) error {
	return c.JSON(http.StatusOK, types.User{Id: 1, Name: "foo"})
}
