package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/info", getInfo)
	r.Run(":8080")
}



func getInfo(c *gin.Context) {
	slackName := c.Query("slack_name")
	track := c.Query("track")

	// Get the current day of the week
	currentDay := time.Now().Weekday().String()

	// Get the current UTC time with validation of +/-2 hours
	currentTime := time.Now().UTC()
	if currentTime.Hour() > 2 && currentTime.Hour() < 22 {
		// Do something if the UTC time is within +/-2 hours
	} else {
		// Handle the case where the UTC time is outside the allowed range
		c.JSON(http.StatusBadRequest, gin.H{"error": "UTC time is outside the allowed range"})
		return
	}

	response := gin.H{
		"slack_name":  slackName,
		"current_day": currentDay,
		"utc_time":    currentTime.Format("2006-01-02T15:04:05Z"),
		"track":       track,
		"status_code": http.StatusOK,
	}
	c.JSON(http.StatusOK, response)
}
