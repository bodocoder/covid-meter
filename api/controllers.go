package api

import (
	"net/http"
	"time"

	"github.com/bodocoder/covid-meter/pkg/db"
	"github.com/bodocoder/covid-meter/pkg/util"
	"github.com/labstack/echo/v4"
)

// this is a controller for route /api/covid?lat=&lon=
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
