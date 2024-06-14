package handler

import (
	"fmt"
	"medication-notifier/crypto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthLogin(ctx *gin.Context) {
	var req LoginRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		panic(fmt.Sprintf("login body err: %s", err))
	}

	// TODO: check user & password
	/////

	authToken, refreshToken, tokenErr := generateTokens("1")
	if tokenErr != nil {
		panic(fmt.Sprintf("login generate token err: %s", tokenErr))
	}
	// TODO: save refresh_token in redis-like storage (with TTL)

	ctx.JSON(http.StatusOK, LoginResponse{
		authToken,
		refreshToken,
	})
}

func AuthRefreshToken(ctx *gin.Context) {
	var req RefreshTokenRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		panic(fmt.Sprintf("refresh_token body err: %s", err))
	}

	userId, err := crypto.ValidateTokenAndReturnUserId(req.RefreshToken)
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// TODO: check with redis-like storage
	authToken, refreshToken, tokenErr := generateTokens(userId)
	if tokenErr != nil {
		panic(fmt.Sprintf("refresh_token generate token err: %s", tokenErr))
	}

	// TODO: revoke previous and save refresh_token in redis-like storage (with TTL)
	ctx.JSON(http.StatusOK, RefreshTokenResponse {
		authToken,
		refreshToken,
	})
}

func generateTokens(userId string) (string, string, error) {
	authToken, err := crypto.GenereteToken(userId, 5) // 5 min
	if err != nil {
		return "", "", err
	}
	refreshToken, err := crypto.GenereteToken(userId, 7 * 60 * 24) // 7 days

	return authToken, refreshToken, err
}

type LoginRequest struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AuthToken   string `json:"auth_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refreh_token"`
}

type RefreshTokenResponse struct {
	AuthToken   string `json:"auth_token"`
	RefreshToken string `json:"refresh_token"`
}
