package util

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/bodocoder/covid-meter/types"
)

func GetCovidCasesIn() map[string]types.CovidCaseState {
	baseUrl, _ := url.Parse("https://data.covid19india.org")
	baseUrl.Path += "v4/min/data.min.json"
	params := url.Values{}
	baseUrl.RawQuery = params.Encode()
	req, _ := http.NewRequest("GET", baseUrl.String(), nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var stateTotalCase map[string]types.CovidCaseState
	_ = json.Unmarshal(body, &stateTotalCase)

	return stateTotalCase
}
