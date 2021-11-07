package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func StartServer() {
	e := echo.New()
	e.GET("/api/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "server is up!")
	})
	e.GET("/api/covid", GetCovidCaseIn)
	e.Logger.Fatal(e.Start(":1323"))
}
