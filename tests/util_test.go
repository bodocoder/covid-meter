package tests

import (
	"testing"

	"github.com/bodocoder/covid-meter/pkg/util"
)

// Test geo util
func TestGetStateIN(t *testing.T) {
	state, err := util.GetStateIN("24.426377570831878", "85.53812213397545")
	if state != "Jharkhand" || err != nil {
		t.Fatalf(`expected state = Jharkhand, found %q, %v, is error`, state, err)
	}
}

//34.05707951936878, 101.4571437380863 outside india location
func TestGetStateNonIN(t *testing.T) {
	state, err := util.GetStateIN("34.05707951936878", "101.4571437380863")
	if state != "" || err == nil {
		t.Fatalf(`expected state = "", found %q, %v, is error`, state, err)
	}
}

//invalid coordinate
func TestGetStateInvalidLatIN(t *testing.T) {
	state, err := util.GetStateIN("", "101.4571437380863")
	if state != "" || err == nil {
		t.Fatalf(`expected state = "", found %q, %v, is error`, state, err)
	}
}

// covid util test

func TestCovidUtil(t *testing.T) {
	covidCaseState := util.GetCovidCasesIn()
	if len(covidCaseState) != 37 {
		t.Fatalf(`expected data of 37 states + total found %v states data`, len(covidCaseState))
	}
}

//33.26955200354856, 130.03380921001252 //saga japan
//21.24957306924999, 81.60505413039007 // nitrr
//24.623584864994612, 85.16992883065895 //vill bihar
