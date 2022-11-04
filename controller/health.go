package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Health struct {
}

func (h *Health) RouteSetup(g *gin.Engine) {
	x := g.Group("health")
	x.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, "OK")
	})
}
