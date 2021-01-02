package main

import (
	"go-blog/handler"
	"go-blog/repository"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" // Using MySQL driver
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/go-playground/validator.v9"
)

var authUser = os.Getenv("AUTH_USER")
var authPassword = os.Getenv("AUTH_PASSWORD")
var db *sqlx.DB
var e = createMax()

func main() {
	db = connectDB()
	repository.SetDB(db)

	auth := e.Group("")

	auth.Use(basicAuth())

	e.GET("/", handler.ArticleIndex)

	e.GET("/articles", handler.ArticleIndex)
	auth.GET("/articles/new", handler.ArticleNew)
	e.GET("/articles/:articleID", handler.ArticleShow)
	auth.GET("/articles/:articleID/edit", handler.ArticleEdit)

	e.GET("/api/articles", handler.ArticleList)
	auth.POST("/api/articles", handler.ArticleCreate)
	auth.DELETE("/api/articles/:articleID", handler.ArticleDelete)
	auth.PATCH("/api/articles/:articleID", handler.ArticleUpdate)

	e.Logger.Fatal(e.Start(":8080"))
}

func connectDB() *sqlx.DB {
	dsn := os.Getenv("DSN")
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		e.Logger.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		e.Logger.Fatal(err)
	}
	log.Println("db connection succeeded")
	return db
}

func createMax() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())
	e.Use(middleware.CSRF())

	e.Static("/css", "src/css")
	e.Static("/js", "src/js")
	e.Validator = &CustomValidator{validator: validator.New()}

	return e
}

// CustomValidator ...
type CustomValidator struct {
	validator *validator.Validate
}

// Validate ...
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

// basicAuth ...
func basicAuth() echo.MiddlewareFunc {
	var basicAuthValidator middleware.BasicAuthValidator

	basicAuthValidator = func(username, password string, c echo.Context) (bool, error) {
		if username == authUser && password == authPassword {
			return true, nil
		}
		return false, nil
	}

	middlewareFunc := middleware.BasicAuth(basicAuthValidator)

	return middlewareFunc
}
