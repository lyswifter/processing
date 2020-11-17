package handler

import "github.com/gin-gonic/gin"

func ShowIndexPage(c *gin.Context) {
	render(c, gin.H{
		"title": "Index",
	}, "index.html")
}
