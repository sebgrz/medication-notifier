package middleware

import (
	"medication-notifier/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ClientInfoMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userAgent := ctx.GetHeader("User-Agent")
		clientId := ctx.GetHeader("X-Client-Id")

		if userAgent == "" || clientId == "" {
			// TODO log
			println("user-agent and client-id headers are required")
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}
		client := utils.ClientInfo{
			Id:   clientId,
			Name: userAgent,
		}

		ctx.Set(utils.CLIENT_INFO_CONTEXT_CONST, client)
	}
}
