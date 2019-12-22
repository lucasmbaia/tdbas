package routes

import (
	"github.com/lucasmbaia/tdbas/api/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

type Group struct {
	groups	[]*gin.RouterGroup
}

func NewGroup(engine *gin.Engine, versions []string) (g Group) {
	for _, v := range versions {
		g.groups = append(g.groups, engine.Group(fmt.Sprintf("/%s", v)))
	}

	return
}

func (g *Group) handle(path string, h gin.HandlerFunc) {
	for _, gr := range g.groups {
		gr.Handle(http.MethodGet, path, h)
		gr.Handle(http.MethodGet, fmt.Sprintf("%s/:ID", path), h)
		gr.Handle(http.MethodPost, path, h)
		gr.Handle(http.MethodDelete, fmt.Sprintf("%s/:ID", path), h)
	}
}

func NewRoutes(engine *gin.Engine, versions []string) {
	var g = NewGroup(engine, versions)

	g.handle("/teste", controllers.NewOrganizations().Relay)

	return
}
