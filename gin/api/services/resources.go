package services

import (
	"github.com/lucasmbaia/tdbas/gin/api/models/interfaces"
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

type ServicesResource struct {
	GetFields func() interface{}
	Model	  func() interfaces.Models
}

func NewServicesResource(f func() interface{}, m func() interfaces.Models) *ServicesResource {
	return &ServicesResource{
		GetFields:  f,
		Model:	    m,
	}
}

func (s *ServicesResource) Get(c *gin.Context) {
	var (
		data  =	s.GetFields()
		m     = s.Model()
	)

	fmt.Println("GET SERVICES")
	data, _ = m.Get(data)

	c.JSON(http.StatusOK, gin.H{"data": data})
	return
}
