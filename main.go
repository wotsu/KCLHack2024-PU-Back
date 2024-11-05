package main

import (
	"net/http"

	"KCLHack-PU-Back/database"
	"KCLHack-PU-Back/services"

	"github.com/labstack/echo/v4"
)

func connect(c echo.Context) error {
	db, _ := database.DB.DB()
	defer db.Close()
	err := db.Ping()
	if err != nil {
		return c.String(http.StatusInternalServerError, "DB接続失敗しました")
	} else {
		return c.String(http.StatusOK, "DB接続しました")
	}
}

func main() {
	e := echo.New()

	// POST
	e.POST("/create/user", services.NewUser)
	e.POST("/create/post", services.NewPost)

	// GET
	e.GET("/", connect)
	e.GET("/get/users", services.GetUsers)
	e.GET("get/posts", services.GetPosts)

	// PUT
	e.PUT("/update/username/:id", services.UpdateUser)

	// DELETE
	e.DELETE("/delete/user/:id", services.DeleteUser)

	e.Logger.Fatal(e.Start(":8080"))
}