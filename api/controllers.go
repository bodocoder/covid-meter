package api

import (
	"net/http"
	"time"

	"github.com/bodocoder/covid-meter/pkg/db"
	"github.com/bodocoder/covid-meter/pkg/util"
	"github.com/labstack/echo/v4"
)

// @Summary get covid case in a state.
// @Description It takes latitude and longitude of a location in India
// @Description and return total covid case in state of that location
// @Tags MainEndpoint
// @Accept json
// @Produce json
// @Param lat query string true "lattitude"
// @Param lon query string true "longitude"
// @Success 200 {object} map[string]interface{}
// @Router /covid [get]
func GetCovidCaseIn(c echo.Context) error {
	lat := c.QueryParam("lat")
	lon := c.QueryParam("lon")
	state, _ := util.GetStateIN(lat, lon) // get the state in india for given lat and lon

	timeUpdated := db.GetUpdateTime()
	if util.AbsTimeDiffInH(time.Now(), timeUpdated) > 1 {
		db.Update() // need an update if covid data in db is older than 1 Hour
	}

	res := db.GetCase(state) // get data from db provided state name
	return c.JSON(http.StatusOK, res)
}

// @Summary show the status of server.
// @Description expected server is up!
// @Tags ServerHealth
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /health [get]
func GetServerHealth(c echo.Context) error {
	return c.JSON(http.StatusOK, "server is up!")
}
