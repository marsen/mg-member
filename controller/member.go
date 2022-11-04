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
	}).PUT(":member_id", func(c *gin.Context) {
		mid := c.Param("member_id")
		c.JSON(http.StatusOK, gin.H{
			"action": fmt.Sprintf("put(replace all) Member %s", mid),
		})
	}).PATCH(":member_id", func(c *gin.Context) {
		mid := c.Param("member_id")
		c.JSON(http.StatusOK, gin.H{
			"action": fmt.Sprintf("patch(update partical) Member %s", mid),
		})
	}).GET(":member_id", func(c *gin.Context) {
		mid := c.Param("member_id")
		c.JSON(http.StatusOK, gin.H{
			"action": fmt.Sprintf("get member %s", mid),
		})
	}).DELETE(":member_id", func(c *gin.Context) {
		mid := c.Param("member_id")
		c.JSON(http.StatusOK, gin.H{
			"action": fmt.Sprintf("delete member %s", mid),
		})
	})
}
