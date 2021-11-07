package util

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

// this does as reverse geo takes latitude and longitude
// return state in india for others error
func GetStateIN(lat string, lon string) (string, error) {
	baseUrl, _ := url.Parse("https://us1.locationiq.com")
	baseUrl.Path += "v1/reverse.php"
	accessKey := os.Getenv("ACCESS_KEY_IQ")
	params := url.Values{}
	params.Add("key", accessKey)
	params.Add("lat", lat)
	params.Add("lon", lon)
	params.Add("format", "json")

	baseUrl.RawQuery = params.Encode()
	req, _ := http.NewRequest("GET", baseUrl.String(), nil)
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	if res.StatusCode != 200 {
		return "", errors.New("not a valid coordinate or location outside India")
	}
	var addressGeo interface{}
	_ = json.Unmarshal(body, &addressGeo)

	address := addressGeo.(map[string]interface{})["address"].(map[string]interface{})
	var country string = address["country"].(string)
	var state string = address["state"].(string)

	if country == "India" {
		return state, nil
	}

	return "", errors.New("not a valid coordinate or location outside India")
}
