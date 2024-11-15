package router

import (
	"github.com/gin-gonic/gin"
)

var Routers []func(*gin.RouterGroup)
