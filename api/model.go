package api

import (
	"database/sql"
	"encoding/json"
)

type Organizations struct {
	ID	    string
	Name	    string
	Description *string
}

type DatabasesTdbas struct {
	ID	      string	      `json:",omitempty"`
	Organization  string	      `json:",omitempty"`
	Name	      string	      `json:",omitempty"`
	Description   string	      `json:",omitempty"`
	Replicas      int	      `json:",omitempty"`
	Status	      string	      `json:",omitempty"`
	Error	      NullString      `json:",omitempty"`
	Username      string	      `json:",omitempty" gorm:"-"`
	Password      string	      `json:",omitempty" gorm:"-"`
}

func (DatabasesTdbas) TableName() string {
	return "databases_tdbas"
}

type HAProxy struct {
	Projects  []Projects  `json:",omitempty"`
}

type Projects struct {
	Name	  string      `json:",omitempty"`
	Databases []Databases `json:",omitempty"`
	Sources	  []string    `json:",omitempty"`
}

type Databases struct {
	Name	    string	  `json:",omitempty"`
	Containers  []Containers  `json:",omitempty"`
}

type Containers struct {
	Minion	string		    `json:",omitempty"`
	Name	string		    `json:",omitempty"`
	Address	string		    `json:",omitempty"`
	Ports	map[string][]string `json:",omitempty"`
}

type NullString struct {
	sql.NullString
}

func (ns *NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(ns.String)
}
