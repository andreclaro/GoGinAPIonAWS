package main

import (
	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ack",
	})
}

func pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": c.Query("name"),
	})
}

func calculator(c *gin.Context) {
	valueA, errA := convertToInteger(c.Query("a"))
	valueB, errB := convertToInteger(c.Query("b"))
	operator := c.Query("operator")
	if errA != nil || errB != nil {
		c.JSON(400, gin.H{
			"message": "Please pass query parameter a, b and operator (+, -, *, /)",
		})
	} else {
		c.JSON(200, gin.H{
			"result": calculate(valueA, valueB, operator),
		})
	}
}
