package interfaces

import (
	"github.com/gin-gonic/gin"
)

type Services interface {
	Get(*gin.Context)
}
