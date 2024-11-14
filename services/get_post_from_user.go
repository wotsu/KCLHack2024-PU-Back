package services

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/golang-jwt/jwt/v5"

	"KCLHack-PU-Back/database"
	"KCLHack-PU-Back/crud"
)

func GetPostsFromUser(c echo.Context) error {

	// すべての投稿を取得
	posts := []database.Post{}
	posts = crud.FetchPosts(posts)

	// ログインしているユーザ名を取得
	user := c.Get("user").(*jwt.Token)
    claims := user.Claims.(jwt.MapClaims)
    userName := claims["name"].(string)

	// ログインしているユーザーの投稿のみを取得
	loginuser_posts := []database.Post{};
	for _, post := range posts {
		if post.UserName == userName {
			loginuser_posts = append(loginuser_posts, post)
		}
	}

	return c.JSON(http.StatusOK, loginuser_posts)

}