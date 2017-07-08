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
	// fakeHeartBeatData()
	// testStrengthInsert()
	// testBatchStrengthInsert()
	// testStrengthQueryOne()
	testStrengthQueryMore()
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

func testQueryMore() {
	sql := "SELECT * FROM heartbeat GROUP BY id LIMIT 10"
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

func testStrengthInsert() {
	sql := `INSERT INTO heartstrength 
		SET timestamp=?, heartstrength=?;`
	_, err := mc.Insert(sql, "heartstrength", 1499522210946, 55)
	if err != nil {
		logrus.Fatalf("insert data to db error: %v", err)
	}
}

func testBatchStrengthInsert() {
	sql := `INSERT INTO heartstrength 
		SET timestamp=?, heartstrength=?;`
	values := []int16{22, 90, 88, 44, 0, 123, 40, 88}
	for _, rate := range values {
		_, err := mc.Insert(sql, "heartstrength", getTimestamp(), rate)
		if err != nil {
			logrus.Fatalf("insert data to db error: %v", err)
		}
	}
}

func testStrengthQueryOne() {
	sql := "SELECT * FROM heartstrength"
	row := mc.QueryOne(sql, "heartstrength")

	var strength entity.HeartStrength
	row.Scan(&strength.ID, &strength.Timestamp, &strength.HeartStrength)

	fmt.Println(strength)
}

func testStrengthQueryMore() {
	sql := "SELECT * FROM heartstrength GROUP BY id LIMIT 10"
	rows, err := mc.QueryMore(sql)
	defer rows.Close()
	if err != nil {
		logrus.Fatal(err)
	}

	var sList []entity.HeartStrength
	for rows.Next() {
		var s entity.HeartStrength
		err := rows.Scan(&s.ID, &s.Timestamp, &s.HeartStrength)
		if err != nil {
			logrus.Fatal(err)
		}

		sList = append(sList, s)
	}

	fmt.Println(sList)
}

// get timestamp whose length is 13
func getTimestamp() int64 {
	return time.Now().UnixNano() / 1000000
}
