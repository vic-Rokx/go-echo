package middleware

import (
	"errors"
	"example/api_gateway/models"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func readCookie(ctx echo.Context) (*http.Cookie, error) {
	cookie, err := ctx.Cookie("Authorization")

	if err != nil {
		return nil, errors.New("could not read cookie")
	}

	return cookie, nil

}

func VerifyAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var user models.User

		cookie, err := readCookie(ctx)

		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		tokenString := cookie.Value
		// Parse takes the token string and a function for looking up the key. The latter is especially
		// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
		// head of the token to identify which key to use, but the parsed token (head and claims) is provided
		// to the callback, providing flexibility.
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return []byte(os.Getenv("SECRET")), nil
		})
		if err != nil {
			log.Fatal(err)
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			currentTime := time.Now().Unix()
			exp := claims["exp"].(float64)

			if float64(currentTime) > exp {
				ctx.String(http.StatusInternalServerError, "Expired token")
			}

			id := claims["sub"]

			// Assuming `db` is your *gorm.DB instance
			result := models.DB.Preload("UserCategories").First(&user, "id = ?", id)

			if result.Error != nil {
				if errors.Is(result.Error, gorm.ErrRecordNotFound) {
					return ctx.String(http.StatusNotExtended, "User not found")
				} else {
					return ctx.String(http.StatusNotExtended, "User not found")
				}
			}

			ctx.Set("user", user)
			return next(ctx)

		} else {
			return ctx.String(http.StatusInternalServerError, "Could not aquire claims, invalid token")
		}
	}
}
