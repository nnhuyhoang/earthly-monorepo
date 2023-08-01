package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nnhuyhoang/earthly-monorepo/libs/hello"
)

func main() {
	e := echo.New()
	e.GET("/one/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, hello.Greet("Tim"))
	})
	_ = e.Start(":8080")
}
