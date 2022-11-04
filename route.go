package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Route struct {
}

func (r *Route) Setup() {
	g := gin.Default()
	g.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	member := g.Group("member")
	member.POST("", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"action": "create new member",
		})
	})
	member.PUT(":member_id", func(c *gin.Context) {
		mid := c.Param("member_id")
		c.JSON(http.StatusOK, gin.H{
			"action": fmt.Sprintf("put(replace all) member %s", mid),
		})
	})
	member.PATCH(":member_id", func(c *gin.Context) {
		mid := c.Param("member_id")
		c.JSON(http.StatusOK, gin.H{
			"action": fmt.Sprintf("patch(update partical) member %s", mid),
		})
	})
	member.GET(":member_id", func(c *gin.Context) {
		mid := c.Param("member_id")
		c.JSON(http.StatusOK, gin.H{
			"action": fmt.Sprintf("get member %s", mid),
		})
	})
	member.DELETE(":member_id", func(c *gin.Context) {
		mid := c.Param("member_id")
		c.JSON(http.StatusOK, gin.H{
			"action": fmt.Sprintf("delete member %s", mid),
		})
	})
	g.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
