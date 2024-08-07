package handler

import (
	"medication-notifier/data"
	"medication-notifier/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *httpHandler) PushTokenRegistration(ctx *gin.Context) {
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
	clientInfo := clientDataAny.(utils.ClientInfo)

	err := h.pushTokenData.Add(data.PushToken{
		UserId: clientInfo.Id,
		Token:  req.Token,
	})
	if err != nil {
		logErrorAndAbort(ctx, "push_registration failed, save token err: %s", err)
		return
	}

	ctx.Status(http.StatusOK)
}

type PushRegisterRequest struct {
	Token string `json:"token"`
}
