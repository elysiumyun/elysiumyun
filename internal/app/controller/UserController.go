package controller

import (
	"net/http"

	"github.com/elysiumyun/elysium/internal/app/model"
	"github.com/gin-gonic/gin"
)

func UserRegister(ctx *gin.Context) {
	userDataCtx, ok := ctx.Get("user")
	if !ok {
		ctx.JSON(http.StatusInternalServerError, "get ctx prams failed!")
		return
	}
	userData := userDataCtx.(model.UserData)
	ctx.JSON(http.StatusOK, userData)
}
