package router

import (
	"github.com/elysiumyun/elysium/internal/pkg/controller"
	"github.com/gin-gonic/gin"
)

// define root routes
func rootRoutes(engine *gin.Engine) {
	apiRoutes := engine.Group("/")
	apiRoutes.GET("health", controller.HealthChecks)
}

// registry all routes
func SetRouter(engine *gin.Engine) {
	rootRoutes(engine)

	// api v1

}
