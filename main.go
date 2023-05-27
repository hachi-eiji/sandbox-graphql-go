package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"http-go-sandbox/handlers"
	"net/http"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())
	e.Use(middleware.CORSWithConfig(corsConfig()))

	e.GET("/hello", handlers.Hello)
	e.GET("/authors/:id", handlers.FindAuthor)
	e.GET("/authors", handlers.ListAuthors)

	e.Logger.Fatal(e.Start(":8000"))
}

func corsConfig() middleware.CORSConfig {
	return middleware.CORSConfig{
		AllowMethods: []string{http.MethodGet},
	}
}
