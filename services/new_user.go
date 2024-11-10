package services

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"KCLHack-PU-Back/database"
	"KCLHack-PU-Back/crud"
)

func NewUser(c echo.Context) error {


	user := database.User{}
	type body struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}
	obj := body{}

	if err := c.Bind(&obj); err != nil {
		return err;
	}
	user.Name = obj.Name
	user.Password = obj.Password

	if user.Name == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "User name is required.")
	}

	if user.Password == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Password  is required.")
	}

	user = crud.CreateUserDB(user)

	if user.ID == 0 {
		return echo.NewHTTPError(http.StatusConflict, "This username is already exist.")
	}

	return c.JSON(http.StatusCreated, user)


}