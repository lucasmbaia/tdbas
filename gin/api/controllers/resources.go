package controllers

import (
	"github.com/lucasmbaia/tdbas/gin/api/services/interfaces"
	"github.com/gin-gonic/gin"
	"fmt"
)

type Resources struct {
	services  interfaces.Services
}

func (r *Resources) Relay(c *gin.Context) {
	switch c.Request.Method {
	case "GET":
		r.Get(c)
	}
}

func (r *Resources) Get(c *gin.Context) {
	fmt.Println("GET CONTROLLERS")
	r.services.Get(c)
	return
}
