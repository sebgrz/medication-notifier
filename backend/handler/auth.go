package handler

import (
	"fmt"
	"medication-notifier/crypto"
	"medication-notifier/data"
	"medication-notifier/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type HttpHandler struct {
	userData  data.UsersDataService
	tokenData data.TokenDataService
}

func New(userData data.UsersDataService, tokenData data.TokenDataService) *HttpHandler {
	return &HttpHandler{
		userData,
		tokenData,
	}
}

func (h *HttpHandler) AuthLogin(ctx *gin.Context) {
	var req LoginRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		panic(fmt.Sprintf("login body err: %s", err))
	}

	clientInfo := ctx.GetString(utils.CLIENT_INFO_CONTEXT_CONST)

	// check user data
	user, err := h.userData.FindByUsername(req.Username)
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}
	if !crypto.ComparePasswordWithHashedPassword(user.Username, req.Password, user.PasswordHash, int(user.CreatedAt)) {
		ctx.AbortWithStatus(http.StatusForbidden)
		return
	}

	authToken, refreshToken, tokenErr := generateTokens(user.Id)
	if tokenErr != nil {
		panic(fmt.Sprintf("login generate token err: %s", tokenErr))
	}
	// TODO: save refresh_token in redis-like storage (with TTL)
	token := data.Token{
		UserId: user.Id,
		Token: refreshToken,
		ExpirationTime: time.Now().Add(time.Minute * 7 * 60 * 24).Unix(),
		ClientInfo: clientInfo,
		ClientId: "TODO",
	}
	if err := h.tokenData.Add(token); err != nil {
		ctx.AbortWithStatus(http.StatusForbidden)
		return
	}

	ctx.JSON(http.StatusOK, LoginResponse{
		authToken,
		refreshToken,
	})
}

func (*HttpHandler) AuthRefreshToken(ctx *gin.Context) {
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
	ctx.JSON(http.StatusOK, RefreshTokenResponse{
		authToken,
		refreshToken,
	})
}

func generateTokens(userId string) (string, string, error) {
	authToken, err := crypto.GenereteToken(userId, 5) // 5 min
	if err != nil {
		return "", "", err
	}
	refreshToken, err := crypto.GenereteToken(userId, 7*60*24) // 7 days

	return authToken, refreshToken, err
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AuthToken    string `json:"auth_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refreh_token"`
}

type RefreshTokenResponse struct {
	AuthToken    string `json:"auth_token"`
	RefreshToken string `json:"refresh_token"`
}
