package middleware

import (
	"net/http"

	"github.com/elysiumyun/elysium/internal/app/model"
	"github.com/gin-gonic/gin"
)

func Entry() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var userData model.UserData = model.UserData{}

		if err := ctx.ShouldBind(&userData); err != nil {
			ctx.JSON(http.StatusInternalServerError, "prams bind failed!")
			ctx.Abort()
			return
		}
		ctx.Set("user", userData)
		ctx.Next()
	}
}
