package api

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/adolphlwq/flipped/entity"
	"github.com/gin-gonic/gin"
)

func heartbeatData(c *gin.Context) {
	var heartbeat entity.HeartBeat
	if err := c.BindJSON(&heartbeat); err != nil {
		logrus.Warnf("parse heartbeat from http post request error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "parse heartbeat from http post request error",
		})
		return
	}

	// save data to db
	sql := `INSERT INTO heartbeat SET timestamp=?, heartrate=?;`
	_, err := mc.Insert(sql, "heartbeat", heartbeat.Timestamp, heartbeat.HeartRate)
	if err != nil {
		logrus.Warnf("save heartbeat data to db error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "save heartbeat data to db error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

// fetchHeartBeatData url schema:
// get /data/heartbeat?limit=10
func fetchHeartBeatData(c *gin.Context) {
	limit, ok := c.GetQuery("limit")
	if !ok {
		limit = "20"
	}

	sql := `SELECT * FROM heartbeat GROUP BY id DESC LIMIT ` + limit + ";"

	rows, err := mc.QueryMore(sql)
	if err != nil {
		logrus.Warnf("select from heartbeat error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "select from heartbeat error",
		})
		return
	}

	var heartList []entity.HeartBeat
	for rows.Next() {
		var heart entity.HeartBeat
		err := rows.Scan(&heart.ID, &heart.Timestamp, &heart.HeartRate)
		if err != nil {
			logrus.Warnf("scan result of rows error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "select from heartbeat error",
			})
			return
		}
		heartList = append(heartList, heart)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": heartList,
	})
}
