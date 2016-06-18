package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RouteMap struct{}

// Register ... Application Routes
func (r *RouteMap) Register(router *gin.RouterGroup) {
	// Demo route * Replace with your application routes *
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "welcome", map[string]string{
			"title":       "Welcome to Pulse!",
			"description": "A bundle of go libs that make building impactful web apps a snap.",
		})
	})
}
