package main

import (
	"example/api_gateway/controllers"
	"example/api_gateway/initializers"
	"example/api_gateway/middleware"
	"example/api_gateway/models"
	"example/api_gateway/routes"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var db *gorm.DB

func userServiceHandler(w http.ResponseWriter, r *http.Request) {
	// In a real setup, this would call the UserService microservice
	fmt.Fprintf(w, "User Service: Responding to request for %s", r.URL.Path[1:])
}

func productServiceHandler(w http.ResponseWriter, r *http.Request) {
	// In a real setup, this would call the ProductService microservice
	fmt.Fprintf(w, "Product Service: Responding to request for %s", r.URL.Path[1:])
}

var mySigningKey = []byte("secret")

func authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		authHeaderParts := strings.Split(authHeader, " ")
		if len(authHeaderParts) != 2 || authHeaderParts[0] != "Bearer" {
			http.Error(w, "Authorization header format must be Bearer {token}", http.StatusUnauthorized)
			return
		}

		tokenString := authHeaderParts[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return mySigningKey, nil
		})

		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			fmt.Printf("Authenticated user: %v\n", claims["user"])
		} else {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Token is valid
		next.ServeHTTP(w, r)
	})
}

func init() {
	initializers.LoadEnvVariables()
}

func main() {

	// config := models.Config{
	// 	Host:     os.Getenv("DB_HOST"),
	// 	Port:     os.Getenv("DB_PORT"),
	// 	User:     os.Getenv("DB_USER"),
	// 	Password: os.Getenv("DB_PASSWORD"),
	// 	DBName:   os.Getenv("DB_NAME"),
	// 	SSLMode:  os.Getenv("DB_SSLMODE"),
	// }

	models.InitDB()

	e := echo.New()

	e.GET("/users/:id", routes.GetUserByID)
	e.GET("/articles/:id", routes.GetArticleById)
	e.GET("/users", routes.GetUsers)
	e.GET("/categories", routes.GetCategories)
	e.GET("/articles", routes.GetArticles)

	e.POST("/signup", controllers.SignUp)
	e.POST("/login", controllers.Login)
	e.POST("/categories", routes.AddCategory)
	e.POST("/articles", routes.AddArticle)
	e.POST("/wipe", routes.Wipe)
	e.GET("/validate", controllers.Validate, middleware.VerifyAuth)

	// router := gin.Default()

	// router.GET("/users", routes.GetUsers)

	// router.GET("/users/:id", routes.GetUserByID)

	// router.PATCH("/users/:id", routes.PatchUser)

	// router.DELETE("/users/:id", routes.DeleteUser)

	// router.POST("/category", routes.AddCategory)
	// router.POST("/users", routes.AddUser)

	// fmt.Println("Server is running on http://localhost:8080")
	// router.Run("localhost:8080")
	e.Logger.Fatal(e.Start(":8080"))

	// http.Handle("/users/", authenticate(http.HandlerFunc(userServiceHandler)))
	// http.Handle("/products/", authenticate(http.HandlerFunc(productServiceHandler)))

	// log.Fatal(http.ListenAndServe(":8080", nil))

}
