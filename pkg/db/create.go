package db

import (
	"context"
	"log"

	"github.com/bodocoder/covid-meter/pkg/util"
)

func InsertAll() interface{} {
	var data []interface{}
	rawData := util.GetCovidCasesIn()
	for k, v := range rawData {
		data = append(data, util.ComplexToSimple(v, util.States[k]))
	}
	collection := util.GetCollection("test", "covid")
	res, err := collection.InsertMany(context.TODO(), data)
	if err != nil {
		log.Fatal(err)
	}
	collection.Database().Client().Disconnect(context.TODO())
	return res.InsertedIDs
}
