package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func FlippedServer() {
	r := gin.Default()
	r.GET("/health", health)

	r.Run(":9090")
}

func health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "health",
	})
}
