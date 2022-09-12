package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response(ctx *gin.Context, httpStatus int, code int, res interface{}, msg string, status string) {
	ctx.JSON(
		httpStatus,
		gin.H{
			"code":    code,
			"result":  res,
			"message": msg,
			"type":    status,
		},
	)
}

func Success(ctx *gin.Context, res interface{}, msg string) {
	Response(ctx, http.StatusOK, 0, res, msg, "success")
}

func Fail(ctx *gin.Context, res interface{}, msg string) {
	Response(ctx, http.StatusOK, -1, res, msg, "error")
}

func Abnormal(ctx *gin.Context, res interface{}, msg string) {
	Response(ctx, http.StatusInternalServerError, -1, res, msg, "error")
}
