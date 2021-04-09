package web

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tabakazu/golang-webapi-demo/controller"
)

type RoutingSet struct {
	Items controller.Items
}

type server struct {
	*echo.Echo
}

func NewServer(r RoutingSet) *server {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})

	e.GET("/items/:id", r.Items.Show)

	return &server{e}
}

func (s server) ListenAndServe() {
	s.Logger.Fatal(s.Start(":8080"))
}