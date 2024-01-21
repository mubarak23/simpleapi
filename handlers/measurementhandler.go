package handlers

import (
 	"simplefitapi/models"
	"simplefitapi/repositories"
	"net/http"
	"strconv"
	"github.com/labstack/echo/v4"
)


type measurmentForm struct {
	UserID int `json:"user_id" form:"user_id"`
	Weight float64 `json:"weight" form:"weight"`
	Height float64 `json:"height" form:"height"`
	BodyFat float64 `json:"body_fat" form:"body_fat"`
}

func CreateMeasurement (c echo.Context) error {
		form := new(measurmentForm)

		if err := c.Bind(form); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		userMeasurement := model.Measurements {
			UserID: form.UserID,
			Weight: form.Weight,
			Height: form.Height,
			BodyFat: form.BodyFat,
		}

		newUserMeasurement, err := repositories.CreateUserMeasurement(userMeasurement)
		
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusCreated, newUserMeasurement)
}

func HandleUpdateMeasurement (c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	measurement := model.Measurements{}
	c.Bind(&measurement)

	updatedMeasurement, err := repositories.UpdateMeasurement(measurement, idInt)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, updatedMeasurement)
}