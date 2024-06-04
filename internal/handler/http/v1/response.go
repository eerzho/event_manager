package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Message string `json:"message"`
}

func errorRsp(ctx *gin.Context, code int, msg string) {
	ctx.AbortWithStatusJSON(code, errorResponse{msg})
}

type successResponse struct {
	Data interface{} `json:"data"`
}

func successRsp(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, successResponse{data})
}
