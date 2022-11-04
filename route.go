package main

import (
	"github.com/gin-gonic/gin"
	"mg-member/controller"
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
	memberCtrl := &controller.Member{}
	memberCtrl.RouteSetup(g)
	g.RedirectFixedPath = true // case insensitivity
	g.Run()                    // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
