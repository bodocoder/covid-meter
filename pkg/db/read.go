package db

import (
	"context"
	"time"

	"github.com/bodocoder/covid-meter/pkg/util"
	"github.com/bodocoder/covid-meter/types"
	"go.mongodb.org/mongo-driver/bson"
)

func GetCase(state string) types.SimpleCovidCaseState {
	filter := bson.D{{"state", state}}
	collection := util.GetCollection("test", "covid")
	var res types.SimpleCovidCaseState
	_ = collection.FindOne(context.TODO(), filter).Decode(&res)

	collection.Database().Client().Disconnect(context.TODO())
	return res
}

func GetUpdateTime() time.Time {
	collection := util.GetCollection("test", "update-time")
	var updateTime types.Time
	_ = collection.FindOne(context.TODO(), bson.D{{}}).Decode(&updateTime)
	collection.Database().Client().Disconnect(context.TODO())
	return updateTime.Val
}
