package handlers

import (
 	"simplefitapi/models"
	"simplefitapi/repositories"
	"net/http"
	"strconv"
	"github.com/labstack/echo/v4"
)

type UserForm struct {
	Name string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

// type EventForm struct {
// 	PubKey string `json:"pubkey" form:"pubkey"`
// 	Kind string `json:"kind" form:"kind"`
// 	Ptags string[] `json:ptags form:"ptags"`
// 	Etags string[] `json:etags form:"etags"`
// 	Gtags string[] `json:"gtags" form:"gtags"`
// 	Content string `json:"content" form:"content"`
// }


// func SubmitEvent (c echo,Context) error {
// 	form := new(EventForm)

// 	if err := c.Bind(form); err != nil {
// 		return c.JSON(http.StatusBadRequest, err)
// 	}

// 	event := model.EventForm{
// 		PubKey: form.pubkey,
// 		Kind: form.Kind,
// 		Ptags: form.Ptags,
// 		Etags: foorm.Etags,
// 		Gtags: form.Gtags,
// 		Content: form.Content
// 	}

// 	newEvent, err := repositories.AddNostrEvent(newEvent)

// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, err.Error())
// 	}

// 	return c.JSON(http.StatusCreated, newEvent)

// }


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

func HandleUpdateUser (c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	user := model.User{}
	c.Bind(&user)
	updateUser, err := repositories.UpdateUser(user, idInt)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, updateUser)

}