package controllers

import (
	"github.com/lucasmbaia/tdbas/gin/api/models/interfaces"
	"github.com/lucasmbaia/tdbas/gin/api/services"
	"github.com/lucasmbaia/tdbas/gin/api/models"
)

type Organizations struct {
	Resources
}

func NewOrganizations() *Organizations {
	var f = func() interface{} {
		return &models.OrganizationsFields{}
	}

	var m = func() interfaces.Models {
		return models.NewOrganizations()
	}

	return &Organizations{Resources{services.NewServicesResource(f, m)}}
}
