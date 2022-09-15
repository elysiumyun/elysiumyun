package middleware

import (
	"github.com/elysiumyun/elysium/internal/app/model"
	"github.com/elysiumyun/elysium/internal/pkg/response"
	"github.com/gin-gonic/gin"
)

func Entry() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var userData model.UserData = model.UserData{}

		if err := ctx.ShouldBind(&userData); err != nil {
			response.Fail(ctx, gin.H{"err": err}, "Prams bind failed!")
			ctx.Abort()
			return
		}
		ctx.Set("user", userData)
		ctx.Next()
	}
}
