package auth

import (
	"KCLHack-PU-Back/crud"
	"net/http"
	"time"
    "os"
    "log"

    "github.com/golang-jwt/jwt/v5"
    "github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

var jwtSecret []byte

func init() {
    // .envファイルを読み込む
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }

    // JWT_SECRET_KEYを取得
    jwtSecret = []byte(os.Getenv("JWT_SECRET_KEY"))
    if len(jwtSecret) == 0 {
        log.Fatal("JWT_SECRET_KEY is not set in .env file")
    }
}

func Login(c echo.Context) error {

	type body struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	obj := body{}
	if err := c.Bind(&obj); err != nil {
		return err;
	}
	
    input_username := obj.Name
    input_password := obj.Password

    // ユーザーネーム -> idにするやつ
    userID, err := crud.FetchUserID(input_username)
    if err != nil {
        return err
    }

    // パスワードを取得
    realPassword, err := crud.FetchPassword(userID)
    if err != nil {
        return err
    }

    // ユーザー認証のロジックをここに追加
    if input_password == realPassword {
        token := jwt.New(jwt.SigningMethodHS256)

        claims := token.Claims.(jwt.MapClaims)
        claims["name"] = input_username
        claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

        t, err := token.SignedString(jwtSecret)
        if err != nil {
            return err
        }

        return c.JSON(http.StatusOK, echo.Map{
            "token": t,
        })
    }

    return echo.ErrUnauthorized
}

func Restricted(c echo.Context) error {
    
    user := c.Get("user").(*jwt.Token)
    claims := user.Claims.(jwt.MapClaims)
    name := claims["name"].(string)
    return c.String(http.StatusOK, "Welcome "+name+"!")
}