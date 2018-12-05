package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Auth struct {
	Token string `json:"token" binding:"required"`
}

func server() {
	r := gin.Default()

	r.Static("/assets", "./static/assets")
	r.LoadHTMLGlob("./templates/*.tmpl")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Super!!",
		})
	})

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

	go r.Run(":12345")
}
