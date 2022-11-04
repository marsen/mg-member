package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Member struct {
}

func (m Member) RouteSetup(g *gin.Engine) {
	member := g.Group("member")
	member.POST("", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"action": "create new Member",
		})
	})
	member.PUT(":member_id", func(c *gin.Context) {
		mid := c.Param("member_id")
		c.JSON(http.StatusOK, gin.H{
			"action": fmt.Sprintf("put(replace all) Member %s", mid),
		})
	})
	member.PATCH(":member_id", func(c *gin.Context) {
		mid := c.Param("member_id")
		c.JSON(http.StatusOK, gin.H{
			"action": fmt.Sprintf("patch(update partical) Member %s", mid),
		})
	})
	member.GET(":member_id", func(c *gin.Context) {
		mid := c.Param("member_id")
		c.JSON(http.StatusOK, gin.H{
			"action": fmt.Sprintf("get Member %s", mid),
		})
	})
	member.DELETE(":member_id", func(c *gin.Context) {
		mid := c.Param("member_id")
		c.JSON(http.StatusOK, gin.H{
			"action": fmt.Sprintf("delete Member %s", mid),
		})
	})
}
