package api

import (
	crand "crypto/rand"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

func FlippedServer() {
	r := gin.Default()
	r.GET("/health", health)
	r.GET("/fake", fake)
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
