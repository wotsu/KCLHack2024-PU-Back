package services

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"KCLHack-PU-Back/database"
	"KCLHack-PU-Back/crud"
)

func GetPosts(c echo.Context) error {

	posts := []database.Post{}
	posts = crud.FetchPosts(posts)
	return c.JSON(http.StatusOK, posts)

}