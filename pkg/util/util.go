package util

import (
	"math"
	"time"

	"github.com/bodocoder/covid-meter/types"
)

var States map[string]string = map[string]string{
	"BR": "Bihar",
	"CH": "Chandigarh",
	"DL": "Delhi",
	"JK": "Jammu and Kashmir",
	"DN": "Dadra and Nagar Haveli and Daman and Diu",
	"HP": "Himachal Pradesh",
	"KA": "Karnataka",
	"MH": "Maharashtra",
	"MZ": "Mizoram",
	"OR": "Odisha",
	"RJ": "Rajsthan",
	"PY": "Puducherry",
	"AP": "Andhra Pradesh",
	"CT": "Chhattisgarh",
	"NL": "Nagaland",
	"PB": "Punjab",
	"WB": "West Bengal",
	"AR": "Arunachal Pradesh",
	"LA": "Ladakh",
	"LD": "Lakshadweep",
	"ML": "Meghalaya",
	"MN": "Manipur",
	"AN": "Andaman and Nicobar Islands",
	"GA": "Goa",
	"TG": "Telangana",
	"UP": "Uttar Pradesh",
	"UT": "Uttarakhand",
	"JH": "Jharkhand",
	"KL": "Kerala",
	"TT": "Total",
	"AS": "Assam",
	"GJ": "Gujrat",
	"HR": "Haryana",
	"MP": "Madhya Pradesh",
	"SK": "Sikkim",
	"TN": "Tamil Nadu",
	"TR": "Tripura",
}

func ComplexToSimple(raw types.CovidCaseState, state string) types.SimpleCovidCaseState {
	return types.SimpleCovidCaseState{
		state,
		raw.Meta.Last_updated,
		raw.Total.Confirmed,
		raw.Total.Deceased,
		raw.Total.Recovered,
	}
}

func AbsTimeDiffInH(u time.Time, v time.Time) int {
	diff := math.Abs(u.Sub(v).Hours())
	return int(diff)
}
