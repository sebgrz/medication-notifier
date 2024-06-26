package handler

import (
	"medication-notifier/data"
	"medication-notifier/utils/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

type httpHandler struct {
	userData  data.UsersDataService
	tokenData data.TokenDataService
}

func New(userData data.UsersDataService, tokenData data.TokenDataService) *httpHandler {
	return &httpHandler{
		userData,
		tokenData,
	}
}

func logErrorAndAbort(ctx *gin.Context, msg string, args ...any) {
	logger.Error(msg, args...)
	ctx.AbortWithStatus(http.StatusBadRequest)
}
