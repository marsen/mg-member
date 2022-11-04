package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseController interface {
	RouteSetup(g *gin.Engine)
}

type Member struct {
}

func (m Member) RouteSetup(g *gin.Engine) {
	c := g.Group("member")
	c.POST("", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"action": "create new member",
		})
	}).PUT(":member_id", func(c *gin.Context) {
		mid := c.Param("member_id")
		c.JSON(http.StatusOK, gin.H{
			"action": fmt.Sprintf("put(replace all) member %s", mid),
		})
	}).PATCH(":member_id", func(c *gin.Context) {
		mid := c.Param("member_id")
		c.JSON(http.StatusOK, gin.H{
			"action": fmt.Sprintf("patch(update partical) member %s", mid),
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
