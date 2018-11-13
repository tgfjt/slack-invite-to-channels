package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Auth struct {
	Token string `json:"token" binding:"required"`
}

func main() {
	r := gin.Default()
	r.StaticFile("/", "./static/index.html")

	r.POST("/members", func(c *gin.Context) {
		var auth Auth

		if err := c.ShouldBindJSON(&auth); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"members": nil,
		})
	})

	r.Run(":12345")
}
