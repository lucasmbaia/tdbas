package controllers

import (
	"encoding/json"
	"net/http"
	"strings"
	"regexp"

	"github.com/lucasmbaia/tdbas/api/model/interfaces"
	"github.com/gin-gonic/gin"
)

type Resources struct {
	GetModel  func() interfaces.Models
	GetFields func() interface{}
}

func (r *Resources) Relay(c *gin.Context) {
	switch c.Request.Method {
	case http.MethodGet:
		r.Get(c)
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "message": "method not implemented."})
	}
}

func (r *Resources) Get(c *gin.Context) {
	var (
		m	= r.GetModel()
		data	= r.GetFields()
		filters	= r.GetFields()
		err	error
	)

	if err = r.setParams(c, filters); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "message": err.Error()})
		return
	}

	if data, err = m.Find(filters); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
	return
}

func (r Resources) setParams(c *gin.Context, f interface{}) (err error) {
	var (
		rgx	= regexp.MustCompile(`\/(:[^:\/]*)`)
		matches	[]string
		params	= make(map[string]interface{})
		str	string
		body	[]byte
	)

	matches = rgx.FindAllString(c.FullPath(), -1)

	for _, v := range matches {
		str = strings.Replace(v, "/:", "", -1)
		params[str] = c.Param(str)
	}

	if body, err = json.Marshal(params); err != nil {
		return
	}

	if err = json.Unmarshal(body, f); err != nil {
		return
	}

	return
}
