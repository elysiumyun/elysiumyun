package controller

import (
	"net/http"
	"time"

	"github.com/elysiumyun/elysium/pkg/info"
	"github.com/gin-gonic/gin"
)

func HealthChecks(ctx *gin.Context) {
	ctx.JSON(
		http.StatusOK,
		gin.H{
			"service":      info.AppName,
			"healthy":      info.SysInfo.Get("healthy"),
			"timestamp":    time.Now().Format(time.ANSIC),
			"dependencies": []gin.H{{}},
		},
	)
}
