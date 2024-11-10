package main

import (
	"log"
	"net/http"
	"os"

	"KCLHack-PU-Back/auth"
	"KCLHack-PU-Back/database"
	"KCLHack-PU-Back/services"

	"github.com/joho/godotenv"
	"github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func init() {
    // .envファイルを読み込む
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }
}

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

	jwtConfig := echojwt.Config{
        SigningKey: []byte(os.Getenv("JWT_SECRET_KEY")),
    }

	// ログインAPI
	e.POST("/login", auth.Login)

	// JWT認証テスト用API
	r := e.Group("/restricted")
	r.Use(echojwt.WithConfig(jwtConfig))
    r.GET("", auth.Restricted)

	// 認証が必要なAPI ----------------------------------------------------
	authGroup := e.Group("/auth")
	authGroup.Use(echojwt.WithConfig(jwtConfig))

	// POST
	authGroup.POST("/create/post", services.NewPost)

	// PUT
	authGroup.PUT("/update/username/:id", services.UpdateUser)
	authGroup.PUT("/update/post/:postId", services.UpdatePost)

	// DELETE
	authGroup.DELETE("/delete/user/:id", services.DeleteUser)
	authGroup.DELETE("/delete/post/:postId", services.DeletePost)

	// 認証が不要なAPI ----------------------------------------------------
	
	// POST
	e.POST("/create/user", services.NewUser)

	// GET
	e.GET("/", connect)
	e.GET("/get/users", services.GetUsers)
	e.GET("/get/posts", services.GetPosts)

	e.Logger.Fatal(e.Start(":8080"))
}