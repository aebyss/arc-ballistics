package main

import (
	"arc-simulator/internal/api"

	"github.com/gin-gonic/gin"
)

func main() {

	api.InitStore()

	r := gin.Default()

	api.RegisterRoutes(r)

	r.Static("/assets", "./frontend/dist/assets")

	r.StaticFile("/favicon.ico", "./frontend/dist/favicon.ico")

	r.GET("/", func(c *gin.Context) {
		c.File("./frontend/dist/index.html")
	})

	r.NoRoute(func(c *gin.Context) {
		c.File("./frontend/dist/index.html")
	})

	r.Run("0.0.0.0:1337")
}
