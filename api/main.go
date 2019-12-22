package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasmbaia/tdbas/api/config"
	"github.com/lucasmbaia/tdbas/api/routes"
)

func loadConfig() {
	config.EnvConfig = config.Config{
		DBUsername:     "tdbas",
		DBPassword:     "totvs@123",
		DBHost:         "127.0.0.1",
		DBPort:         "3306",
		DBName:		"tdbas",
		DBTimeout:      "30000ms",
		DBDebug:        true,
		DBConnsMaxIdle: 5,
		DBConnsMaxOpen: 5,
	}

	config.LoadDB()
}

func main() {
	var (
		g   *gin.Engine
	)

	loadConfig()

	g = gin.Default()

	routes.NewRoutes(g, []string{"v1"})
	g.Run()
}
