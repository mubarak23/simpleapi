package handlers

import (
 	"simplefitapi/models"
	"simplefitapi/repositories"
	"net/http"
	"github.com/labstack/echo/v4"
)

type UserForm struct {
	Name string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}


func CreateUser(c echo.Context) error {
	
	form := new(UserForm) 

if err := c.Bind(form); err != nil {
	return c.JSON(http.StatusBadRequest, err)
}


	user :=  model.User{
		Name: form.Name,
		Email: form.Email,
		Password: form.Password,
	}

	// we are not hashing the password at the moment - coming back to it
	newUser, err := repositories.CreateUser(user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, newUser)
}

func updateUser (c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	updateUser, err := repositories.UpdateUser(user, idInt)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, updateUser)

}