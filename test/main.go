package main

import (
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/adolphlwq/flipped/storage"
)

var mc *storage.MysqlClient

func main() {
	mc = storage.NewMysqlClient()
	testInsert()
}

func testInsert() {
	heartbeatInsertSql := `INSERT INTO heartbeat 
		SET timestamp=?, heartrate=?;`
	values := []int16{22, 90, 88, 44, 0, 123, 40, 88}
	for _, rate := range values {
		_, err := mc.Insert(heartbeatInsertSql, "hearteat", getTimestamp(), rate)
		if err != nil {
			logrus.Fatalf("insert data to db error: %v", err)
		}
	}
}

func getTimestamp() int64 {
	return time.Now().Unix()
}
