package model

import (
	"github.com/lucasmbaia/tdbas/api/model/interfaces"
)

type Resources struct {
	m     interfaces.Models
}

func NewResources(m interfaces.Models) *Resources {
	return &Resources{m}
}

func (r *Resources) Find(fields interface{}) (interface{}, error) {
	return r.m.Find(fields)
}

func (r *Resources) Create(data interface{}) (bool, error) {
	return r.m.Create(data)
}

func (r *Resources) Remove(fields interface{}) (bool, error) {
	return false, nil
}
