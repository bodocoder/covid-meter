package db

import (
	"context"
	"log"

	"github.com/bodocoder/covid-meter/pkg/util"
)

//
func DeleteAll() {
	collection := util.GetCollection("test", "covid")
	err := collection.Drop(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
}
