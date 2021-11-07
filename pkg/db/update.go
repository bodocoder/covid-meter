package db

import (
	"context"
	"time"

	"github.com/bodocoder/covid-meter/pkg/util"
	"github.com/bodocoder/covid-meter/types"
)

func Update() {
	//1. drop current backup
	//2. create new backup
	//3. drop current
	//4. insert new

	//backup TODO

	DeleteAll()
	InsertAll()
	updateUpdateTime()
}

func updateUpdateTime() {
	collection := util.GetCollection("test", "update-time")
	collection.Drop(context.TODO())
	var curTime types.Time
	curTime.Val = time.Now()
	collection.InsertOne(context.TODO(), curTime)
	collection.Database().Client().Disconnect(context.TODO())
}
