package api

import (
	"github.com/lucasmbaia/tdbas/repository"
	"github.com/lucasmbaia/tdbas/etcd"
	"testing"
)

func configApi() TdbasAPIConfig {
	return TdbasAPIConfig {
		Port: 8080,
		DBConfig: repository.RepositoryConfig{
			Username:     "tdbas",
			Password:     "totvs@123",
			Host:	      "127.0.0.1",
			Port:	      "3306",
			DBName:	      "tdbas",
			Timeout:      "30000ms",
			Debug:	      true,
			ConnsMaxIdle: 5,
			ConnsMaxOpen: 5,
		},
		Etcd:	etcd.Config{
			Endpoints:  []string{"http://127.0.0.1:2379"},
			Timeout:    30,
		},
	}
}

func Test_Start(t *testing.T) {
	var (
		td  TdbasAPI
		err error
	)

	if td, err = NewTdbasAPI(configApi()); err != nil {
		t.Fatal(err)
	}

	td.Start()
}
