package handlers

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

// function accept echo.Context
// echo.Context contains the request and response objects.
func Home (c echo.Context) error {
	return c.String(http.StatusOK, "This is the First Access Point")
}