package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func main()  {
	e := echo.New()

	e.GET("/test", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct {
			Message string `json:"message"`
		}{Message: "success!"})
	})

	e.Logger.Fatal(e.Start(":80"))
}