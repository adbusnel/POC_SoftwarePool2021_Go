package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func world(c *gin.Context) {
	c.String(http.StatusOK, "world")
}

func ApplyRoutes(r *gin.Engine) {
	r.GET("/hello", world)
}
