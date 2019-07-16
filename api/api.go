package api

import (
	"github.com/lucasmbaia/tdbas/repository"
	"github.com/lucasmbaia/tdbas/core"
	"github.com/lucasmbaia/tdbas/utils"
	"github.com/lucasmbaia/tdbas/etcd"
	"github.com/satori/go.uuid"
	"github.com/gin-gonic/gin"
	"database/sql"
	"net/http"
	"context"
	"fmt"
)

const (
	defaultPort = ":8080"
	key	    = "/tdbas-haproxy/"
)

type TdbasAPI struct {
	port  string
	repo  repository.Repository
	g     *gin.Engine
	c     core.Core
	etcd  etcd.Client
	ha    utils.HAProxy
}

type TdbasAPIConfig struct {
	DBConfig  repository.RepositoryConfig
	Port	  int
	Etcd	  etcd.Config
}

func NewTdbasAPI(cfg TdbasAPIConfig) (t TdbasAPI, err error) {
	var ctx = context.Background()

	if cfg.Port == 0 {
		t.port = defaultPort
	} else {
		t.port = fmt.Sprintf(":%d", cfg.Port)
	}

	if t.repo, err = repository.NewRepository(cfg.DBConfig); err != nil {
		return
	}

	if t.c, err = core.NewCore(ctx); err != nil {
		return
	}

	if t.etcd, err = etcd.NewClient(ctx, cfg.Etcd); err != nil {
		return
	}

	t.ha = utils.NewHAProxy(utils.KEY_ETCD, t.etcd)
	t.g = gin.Default()
	return
}

func (t *TdbasAPI) Start() {
	v1 := t.g.Group("/api/v1/organizations")
	{
		v1.GET("", t.fetchOrganizations)
		v1.GET("/:organization_id/databases", t.fetchAllDatabases)
		v1.POST("/:organization_id/databases", t.createDatabase)
	}

	t.g.Run()
}

func (t *TdbasAPI) fetchOrganizations(c *gin.Context) {
	var (
		org []Organizations
		err error
	)

	if _, err = t.repo.ReadAll(&org); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "message": err.Error()})
		return
	}

	if len(org) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "message": "No todo found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": org})
}

func (t *TdbasAPI) fetchAllDatabases(c *gin.Context) {
	var (
		db  []DatabasesTdbas
		err error
	)

	if _, err = t.repo.ReadAll(&db); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "message": err.Error()})
		return
	}

	if len(db) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "message": "No todo found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": db})
}

func (t *TdbasAPI) createDatabase(c *gin.Context) {
	var (
		db  DatabasesTdbas
		err error
		id  uuid.UUID
	)

	if err = c.BindJSON(&db); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "message": err.Error()})
		return
	}

	if id, err = uuid.NewV4(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "message": err.Error()})
		return
	}

	db.Organization = c.Param("organization_id")
	db.ID = id.String()
	db.Status = "IN_PROGRESS"

	if err = t.repo.Create(&db); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "message": err.Error()})
		return
	}

	go t.createContainer(db)
	c.JSON(http.StatusCreated, gin.H{"id": db.ID})
}

func (t *TdbasAPI) createContainer(db DatabasesTdbas) {
	var (
		err   error
		cs    []core.Container
	)

	if cs, err = t.c.CreateDB(core.Container{
		Name:	  db.Name,
		Replicas: db.Replicas,
	}); err != nil {
		t.repo.Update(&DatabasesTdbas{ID: db.ID}, &DatabasesTdbas{Status: "ERROR", Error: NullString{sql.NullString{String: err.Error(), Valid: true}}})
		return
	}

	for _, c := range cs {
		var ports = make(map[string][]string)

		for _, p := range c.Ports {
			ports[p.Source] = p.Destinations
		}

		if err = t.ha.GenerateConf(utils.HAProxyConfig{
			Customer:	  "lucas",
			ApplicationName:  db.Name,
			ContainerName:	  c.Name,
			PortsContainer:	  ports,
			Protocol:	  map[string]string{"1433": "tcp"},
			AddressContainer: c.Address,
			Dns:		  "lucas.local",
			Minion:		  "minion-1",
		}); err != nil {
			t.repo.Update(&DatabasesTdbas{ID: db.ID}, &DatabasesTdbas{Status: "ERROR", Error: NullString{sql.NullString{String: err.Error(), Valid: true}}})
			return
		}
	}

	t.repo.Update(&DatabasesTdbas{ID: db.ID}, &DatabasesTdbas{Status: "COMPLETED"})
}
