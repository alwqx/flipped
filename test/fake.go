package main

import (
	crand "crypto/rand"
	"math/rand"
	"time"

	"github.com/Sirupsen/logrus"
)

func fakeHeartBeatData() {
	hearts := []int16{22, 90, 88, 44, 0, 123, 40, 88, 45, 123, 144, 131}
	sql := `INSERT INTO heartbeat 
		SET timestamp=?, heartrate=?;`
	for {
		rate := hearts[rand.Intn(len(hearts))]
		ts := getTimestamp()
		_, err := mc.Insert(sql, "hearteat", ts, rate)
		if err != nil {
			logrus.Fatalf("insert data to db error: %v", err)
		}
		logrus.Infof("insert timestamp %d and rate %d to db", ts, rate)
		time.Sleep(time.Millisecond * 20)
	}
}

func random() int {
	b := make([]byte, 0)
	_, err := crand.Read(b)
	if err != nil {
		logrus.Fatal(err)
	}

	return int(b[0])
}
