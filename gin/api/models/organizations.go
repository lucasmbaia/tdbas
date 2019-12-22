package models

import (
	"fmt"
)

type OrganizationsFields struct {
	ID  string  `json:",omitempty"`
}

type Organizations struct {
}

func NewOrganizations() *Organizations {
	return &Organizations{}
}

func (o *Organizations) Get(filters interface{}) (interface{}, error) {
	fmt.Println("TAMO NO MODEL")
	var (
		data  = &OrganizationsFields{}
		err   error
	)

	data.ID	= "TESTE"
	return data, err
}

func (o *Organizations) Post(data *OrganizationsFields) {
	return
}
