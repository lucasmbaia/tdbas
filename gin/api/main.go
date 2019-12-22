package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasmbaia/tdbas/gin/api/controllers"
)

func main() {
	var (
		g   *gin.Engine
		org *controllers.Organizations
	)

	org = controllers.NewOrganizations()
	g = gin.Default()

	v1 := g.Group("/teste")
	{
		v1.GET("", org.Relay)
	}

	g.Run()
}
