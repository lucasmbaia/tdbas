package config

import (
	"github.com/lucasmbaia/tdbas/repository"
)

func LoadDB() {
	var err error

	if EnvSingletons.DB, err = repository.NewRepository(repository.RepositoryConfig{
		Username:	  EnvConfig.DBUsername,
		Password:	  EnvConfig.DBPassword,
		Host:		  EnvConfig.DBHost,
		Port:		  EnvConfig.DBPort,
		DBName:		  EnvConfig.DBName,
		Timeout:	  EnvConfig.DBTimeout,
		Debug:		  EnvConfig.DBDebug,
		ConnsMaxIdle:	  EnvConfig.DBConnsMaxIdle,
		ConnsMaxOpen:	  EnvConfig.DBConnsMaxOpen,
		ConnsMaxLifetime: EnvConfig.DBConnsMaxLifetime,
	}); err != nil {
		panic(err)
	}
}
