package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthLogin(ctx *gin.Context) {
	var req LoginRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		panic(fmt.Sprintf("login body err: %s", err))
	}

	ctx.JSON(http.StatusOK, LoginResponse{
		"TODO",
		"TODO",
	})
}

type LoginRequest struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AuthToken   string `json:"auth_token"`
	RefresToken string `json:"refresh_token"`
}
