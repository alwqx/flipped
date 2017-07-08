package main

import (
	"fmt"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/adolphlwq/flipped/entity"
	"github.com/adolphlwq/flipped/storage"
)

var mc *storage.MysqlClient

func main() {
	mc = storage.NewMysqlClient()
	// testInsert()
	// testQueryOne()
	// testQueryMore("1")
	// testQueryMore("-1")
	// testBatchInsert()
	// fmt.Println(getTimestamp())
	fakeHeartBeatData()
}

func testInsert() {
	heartbeatInsertSql := `INSERT INTO heartbeat 
		SET timestamp=?, heartrate=?;`
	_, err := mc.Insert(heartbeatInsertSql, "hearteat", 1499522210946, 55)
	if err != nil {
		logrus.Fatalf("insert data to db error: %v", err)
	}
}

func testBatchInsert() {
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

func testQueryOne() {
	sql := "SELECT * FROM heartbeat"
	row := mc.QueryOne(sql, "heartbeat")

	var heartbeat entity.HeartBeat
	row.Scan(&heartbeat.ID, &heartbeat.Timestamp, &heartbeat.HeartRate)

	fmt.Println(heartbeat)
}

func testQueryMore(sortTag string) {
	tag := "ASC"
	if sortTag == "1" {
		tag = "DESC"
	}
	sql := "SELECT * FROM heartbeat GROUP BY id " + tag + " LIMIT 10"
	rows, err := mc.QueryMore(sql)
	defer rows.Close()
	if err != nil {
		logrus.Fatal(err)
	}

	var heartList []entity.HeartBeat
	for rows.Next() {
		var heart entity.HeartBeat
		err := rows.Scan(&heart.ID, &heart.Timestamp, &heart.HeartRate)
		if err != nil {
			logrus.Fatal(err)
		}

		heartList = append(heartList, heart)
	}

	fmt.Println(heartList)
}

// get timestamp whose length is 13
func getTimestamp() int64 {
	return time.Now().UnixNano() / 1000000
}
