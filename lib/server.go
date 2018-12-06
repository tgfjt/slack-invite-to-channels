package lib

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Auth struct {
	Token string `json:"token" binding:"required"`
}

func Server() {
	r := gin.Default()

	r.Static("/assets", "./public/assets")
	r.LoadHTMLGlob("./templates/*.tmpl")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Welcome to Lacks:)",
		})
	})

	r.POST("/members", func(c *gin.Context) {
		var auth Auth

		if err := c.ShouldBindJSON(&auth); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var members Members
		members = GetMembers(auth.Token)

		c.JSON(200, gin.H{
			"data": members,
		})
	})

	r.POST("/channels", func(c *gin.Context) {
		var auth Auth

		if err := c.ShouldBindJSON(&auth); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var channels ChannelsList
		channels = GetChannels(auth.Token)

		c.JSON(200, gin.H{
			"data": channels,
		})
	})

	r.Run(":12345")
}
