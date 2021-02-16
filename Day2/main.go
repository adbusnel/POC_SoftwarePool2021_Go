package main

import (
	router "SoftwareGoDay2/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.ApplyRoutes(r)
	r.Run()
}
