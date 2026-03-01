package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func currentDate() string {
	return time.Now().Format("2006-01-02")
}

func main() {
	r := gin.Default()

	r.GET("/currentdate", func(c *gin.Context) {
		currentDate := currentDate()

		c.JSON(http.StatusOK, gin.H{
			"date": currentDate,
		})
	})

	// Start server on port 8080 (default)
	if err := r.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
