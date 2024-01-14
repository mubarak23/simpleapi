package handlers

import (
	"simplefitapi/models"
	"simplefitapi/repositories"
	"net/http"
	"github.com/labstack/echo/v4"
)


func CreateUser(c echo.Context) error {
	user :=  model.User{}

	c.Bind(&user)

	// we are not hashing the password at the moment - coming back to it
	newUser, err := repositories.CreateUser(user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, newUser)
}

