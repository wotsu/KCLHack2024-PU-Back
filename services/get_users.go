package services

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"KCLHack-PU-Back/database"
	"KCLHack-PU-Back/crud"
)

func GetUsers(c echo.Context) error {

	users := []database.User{}
	users = crud.FetchUsers(users)
	return c.JSON(http.StatusOK, users)

}