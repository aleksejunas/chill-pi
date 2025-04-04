package routes

import (
	"github.com/gin-gonic/gin"
)

// ResourceHandler defines all CRUD-methods as a handler that can implement
type ResourceHandler interface {
	GetAll(*gin.Context)
	GetOne(*gin.Context)
	Create(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
}

// Registers standard CRUD-endpoints for a given resource (prefix)
func RegisterResourceRoutes(r *gin.Engine, basePath string, h ResourceHandler) {
	group := r.Group(basePath)

	// Without ID
	group.GET("/", h.GetAll)
	group.POST("/", h.Create)

	// With ID
	group.GET("/:id", h.GetOne)
	group.PUT("/:id", h.Update)
	group.DELETE("/:id", h.Delete)
}

// RegisterLooseRoute kobler både /path og /path/ til samme handler
func RegisterLooseRoute(group *gin.RouterGroup, method string, path string, handler gin.HandlerFunc) {
	// uten trailing slash
	switch method {
	case "GET":
		group.GET(path, handler)
		group.GET(path+"/", handler)
	case "POST":
		group.POST(path, handler)
		group.POST(path+"/", handler)
	case "PUT":
		group.PUT(path, handler)
		group.PUT(path+"/", handler)
	case "DELETE":
		group.DELETE(path, handler)
		group.DELETE(path+"/", handler)
		// legg til flere metoder hvis du trenger
	}
}
