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
		VALUES (?, ?)`

	_, err := mc.Insert(heartbeatInsertSql, "hearteat", getTimestamp(), 23.4)
	if err != nil {
		logrus.Fatalf("insert data to db error: %v", err)
	}
}

func getTimestamp() int64 {
	return time.Now().Unix()
}
