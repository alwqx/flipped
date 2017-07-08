package api

import (
	crand "crypto/rand"
	"net/http"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/adolphlwq/flipped/storage"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var mc *storage.MysqlClient

func FlippedServer() {
	mc = storage.NewMysqlClient()
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowMethods: []string{"*"},
		AllowHeaders: []string{"Content-Type"},
		ExposeHeaders: []string{"Content-Length", "Access-Control-Allow-Origin",
			"Access-Control-Allow-Headers", "Access-Control-Allow-Methods"},
		//AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			logrus.Info("origin is ", origin)
			return true
		},
		MaxAge: 12 * time.Hour,
	}))

	r.GET("/health", health)
	r.GET("/fake", fake)
	r.POST("/data/heartbeat", heartbeatData)
	r.GET("/data/heartbeat", fetchHeartBeatData)

	r.Run(":9090")
}

func health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "health",
	})
}

func fake(c *gin.Context) {
	b := make([]byte, 1)
	_, err := crand.Read(b)
	if err != nil {
		logrus.Fatalf("generate fake data error: %v", err)
	}
	c.JSON(http.StatusOK, gin.H{"data": b[0]})
}
