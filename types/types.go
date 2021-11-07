package types

import "time"

type Total struct {
	Confirmed int
	Deceased  int
	Recovered int
}

type Meta struct {
	Last_updated string
}

type CovidCaseState struct {
	Meta  Meta
	Total Total
}

type SimpleCovidCaseState struct {
	State       string
	LastUpdated string
	Confirmed   int
	Deceased    int
	Recovered   int
}

type Time struct {
	Val time.Time
}
