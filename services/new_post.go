package services

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/golang-jwt/jwt/v5"

	"KCLHack-PU-Back/database"
	"KCLHack-PU-Back/crud"
)

func NewPost(c echo.Context) error {

	post := database.Post{}
	type body struct {   
		Task       string    `json:"task"`
		Importance uint      `json:"importance"`
		Deadline   string    `json:"deadline"`
	}

	obj := body{}
	if err := c.Bind(&obj); err != nil {
		return err;
	}

	user := c.Get("user").(*jwt.Token)
    claims := user.Claims.(jwt.MapClaims)
    name := claims["name"].(string)

	post.UserName = name
	post.Task = obj.Task
	post.Importance = obj.Importance
	post.Deadline = obj.Deadline


	createedPost, err := crud.CreatePostDB(post)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to create post")
	}

	if createedPost.ID == 0 {
		return echo.NewHTTPError(http.StatusConflict, "failed to create post | id error")
	}

	return c.JSON(http.StatusCreated, createedPost)

}