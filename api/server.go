package api

import (
	_ "github.com/bodocoder/covid-meter/docs"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func StartServer() {
	e := echo.New()
	apiGroup := e.Group("/api/v1")

	apiGroup.GET("/health", GetServerHealth)
	apiGroup.GET("/covid", GetCovidCaseIn)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Logger.Fatal(e.Start(":1323"))
}
