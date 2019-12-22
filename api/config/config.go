package config

import (
	"github.com/lucasmbaia/tdbas/repository"
)

var (
	EnvConfig     Config
	EnvSingletons Singletons
)

type Config struct {
	DBUsername	    string  `json:",omitempty"`
	DBPassword	    string  `json:",omitempty"`
	DBHost		    string  `json:",omitempty"`
	DBPort		    string  `json:",omitempty"`
	DBName		    string  `json:",omitempty"`
	DBTimeout	    string  `json:",omitempty"`
	DBDebug		    bool    `json:",omitempty"`
	DBConnsMaxIdle	    int	    `json:",omitempty"`
	DBConnsMaxOpen	    int	    `json:",omitempty"`
	DBConnsMaxLifetime  int	    `json:",omitempty"`
}

type Singletons struct {
	DB  repository.Repository
}
