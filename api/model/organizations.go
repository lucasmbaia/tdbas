package model

import (
	"github.com/lucasmbaia/tdbas/repository"
	"github.com/lucasmbaia/tdbas/api/datamodels"
)

/*type OrganizationsFields struct {
	ID	    string  `json:",omitempty"`
	Name	    string  `json:",omitempty"`
	Description *string `json:",omitempty"`
}

func (OrganizationsFields) TableName() string {
	return "organizations"
}*/

type Organizations struct {
	Resources

	repo  repository.Repository
}

func NewOrganizations(r	repository.Repository) *Organizations {
	return &Organizations{repo: r}
}

func (o *Organizations) Find(filters interface{}) (interface{}, error) {
	var (
		entity	= []datamodels.Organizations{}
		err	error
	)

	if _, err = o.repo.ReadAll(filters, &entity); err != nil {
		return nil, err
	}

	return entity, err
}

func (o *Organizations) Create(data interface{}) (async bool, err error) {
	return
}
