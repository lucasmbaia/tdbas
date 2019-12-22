package main

import (
	"github.com/lucasmbaia/tdbas/etcd"
	"github.com/lucasmbaia/tdbas/repository"
	"github.com/lucasmbaia/tdbas/api"
)

func main() {
	var (
		td  api.TdbasAPI
		err error
	)

	if td, err = api.NewTdbasAPI(api.TdbasAPIConfig{
		Port: 8080,
		DBConfig: repository.RepositoryConfig{
			Username:     "tdbas",
			Password:     "totvs@123",
			Host:         "127.0.0.1",
			Port:         "3306",
			DBName:       "tdbas",
			Timeout:      "30000ms",
			Debug:        true,
			ConnsMaxIdle: 5,
			ConnsMaxOpen: 5,
		},
		Etcd:   etcd.Config{
			Endpoints:  []string{"http://172.16.95.183:2379"},
			Timeout:    30,
		},
	}); err != nil {
		panic(err)
	}

	td.Start()
}
