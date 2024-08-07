package handler

import (
	"github.com/gin-gonic/gin"
	"medication-notifier/utils"
)

func (h *httpHandler) PushRegistration(ctx *gin.Context) {
	var req PushRegisterRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		logErrorAndAbort(ctx, "push_registration body err: %s", err)
		return
	}
	var clientDataAny, exists = ctx.Get(utils.CLIENT_INFO_CONTEXT_CONST)
	if !exists {
		logErrorAndAbort(ctx, "push_registration failed, clientData is empty")
		return
	}
	_ = clientDataAny.(utils.ClientInfo)
}

type PushRegisterRequest struct {
	Token string `json:"token"`
}
