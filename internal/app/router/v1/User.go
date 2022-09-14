package v1

import (
	"github.com/elysiumyun/elysium/internal/app/controller"
	"github.com/elysiumyun/elysium/internal/app/middleware"
	"github.com/gin-gonic/gin"
)

func User(engine *gin.Engine) {
	apiRoutes := engine.Group(prefix + "/user")
	apiRoutes.POST("/register", middleware.Entry(), controller.UserRegister)
}
