package main

import (
	"net/http"
	"time"

	"github.com/flosch/pongo2"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

const tmplPath = "src/template/"

var e = createMax()

func main() {
	e.GET("/", articleIndex)

	e.Logger.Fatal(e.Start(":8080"))
}

func createMax() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	e.Static("/css", "src/css")
	e.Static("/js", "src/js")

	return e
}

func articleIndex(c echo.Context) error {
	data := map[string]interface{}{
		"Message": "Hello, World!",
		"Now":     time.Now(),
	}
	return render(c, "article/index.html", data)
}

func htmlBlob(file string, data map[string]interface{}) ([]byte, error) {
	return pongo2.Must(pongo2.FromCache(tmplPath + file)).ExecuteBytes(data)
}

func render(c echo.Context, file string, data map[string]interface{}) error {
	b, err := htmlBlob(file, data)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.HTMLBlob(http.StatusOK, b)
}
