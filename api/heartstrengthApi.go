package api

import (
	"fmt"
	"net/http"

	"strconv"

	"github.com/Sirupsen/logrus"
	"github.com/adolphlwq/flipped/entity"
	"github.com/gin-gonic/gin"
)

// postHeartStrengthData post /data/heartstrength/:batchsize
func postHeartStrengthData(c *gin.Context) {
	batchSizeStr := c.Param("batchsize")
	if len(batchSizeStr) == 0 {
		logrus.Warnf("batchsize error")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "need correct batch size",
		})

		return
	}
	batchSize, err := strconv.ParseInt(batchSizeStr, 10, 64)
	if err != nil {
		logrus.Warnf("convert batchsize error")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "convert batchsize error",
		})

		return
	}

	var hsm entity.HSM
	if err := c.BindJSON(&hsm); err != nil {
		logrus.Warnf("parse key heartstrength from http post request error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "parse key heartstrength from http post request error",
		})
		return
	}

	// save data to db
	for i := 0; int64(i) < batchSize; i++ {
		key := strconv.FormatInt(int64(i), 10)
		heartStrength := hsm[key]
		sql := `INSERT INTO heartstrength SET timestamp=?, heartstrength=?;`
		_, err := mc.Insert(sql, "heartstrength", heartStrength.Timestamp, heartStrength.HeartStrength)
		if err != nil {
			logrus.Warnf("insert heart strength data to db error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "insert heart strength data to db error",
			})
			return
		}
	}

	// dave data of arrays to db success
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("save all data of %d datas to db", batchSize),
	})
}

// fetchHeartStrengthData url schema:
// get /data/heartbeat?limit=10
func fetchHeartStrengthData(c *gin.Context) {
	limit, ok := c.GetQuery("limit")
	if !ok {
		limit = "20"
	}

	sql := `SELECT * FROM heartstrength GROUP BY id DESC LIMIT ` + limit + ";"

	rows, err := mc.QueryMore(sql)
	if err != nil {
		logrus.Warnf("select from heartstrength error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "select from heartstrength error",
		})
		return
	}

	var strengthList []entity.HeartStrength
	for rows.Next() {
		var strength entity.HeartStrength
		err := rows.Scan(&strength.ID, &strength.Timestamp, &strength.HeartStrength)
		if err != nil {
			logrus.Warnf("scan result of rows error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "select from heartstrength error",
			})
			return
		}
		strengthList = append(strengthList, strength)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": strengthList,
	})
}

// strAddInt("11", 3) => "13"
func strAddInt(key string, toAdd int16) string {
	keyInt, err := strconv.ParseInt(key, 10, 64)
	if err != nil {
		logrus.Fatalf("parse string to int error %v", err)
	}
	sum := int16(keyInt) + toAdd
	sumStr := strconv.FormatInt(int64(sum), 10)
	return sumStr
}

func intAddStr(key int16, toAdd string) string {
	addInt, err := strconv.ParseInt(toAdd, 10, 64)
	if err != nil {
		logrus.Fatalf("parse int to string error: %v", err)
	}
	sum := key + int16(addInt)

	return strconv.FormatInt(int64(sum), 10)
}
