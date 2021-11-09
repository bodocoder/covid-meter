package main

import (
	"github.com/bodocoder/covid-meter/api"
)

// @title Covid-Meter
// @version 1.0
// @description API to fetch covid case in a state of India.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name license.name Yet to get
// @license.url license.url.yet.to.get

// @host localhost:1323
// @BasePath /api/v1
func main() {
	api.StartServer()
}
