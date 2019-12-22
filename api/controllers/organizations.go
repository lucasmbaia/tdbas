package controllers

import (
	"github.com/lucasmbaia/tdbas/api/model/interfaces"
	"github.com/lucasmbaia/tdbas/api/config"
	"github.com/lucasmbaia/tdbas/api/model"
	"github.com/lucasmbaia/tdbas/api/datamodels"
)

type Organizations struct {
	Resources
}

func NewOrganizations() *Organizations {
	return &Organizations{
		Resources{
			GetModel:   func() interfaces.Models {
				return model.NewResources(model.NewOrganizations(config.EnvSingletons.DB))
			},
			GetFields:  func() interface{} {
				return &datamodels.Organizations{}
			},
		},
	}
}
