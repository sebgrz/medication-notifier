package handler

import (
	"medication-notifier/crypto"
	"medication-notifier/data"
	"medication-notifier/utils"
	"medication-notifier/utils/logger"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *httpHandler) AuthLogin(ctx *gin.Context) {
	var req LoginRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		logErrorAndAbort(ctx, "login body err: %s", err)
		return
	}

	clientInfoRaw, _ := ctx.Get(utils.CLIENT_INFO_CONTEXT_CONST)
	clientInfo := clientInfoRaw.(utils.ClientInfo)

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

	authToken, refreshToken, expAt, tokenErr := generateTokens(user.Id)
	if tokenErr != nil {
		logErrorAndAbort(ctx, "login generate token err: %s", tokenErr)
		return
	}

	// save refresh_token in temporary storage
	token := data.Token{
		UserId:         user.Id,
		Token:          refreshToken,
		ExpirationTime: expAt,
		ClientInfo:     clientInfo.Name,
		ClientId:       clientInfo.Id,
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

func (h *httpHandler) AuthRefreshToken(ctx *gin.Context) {
	var req RefreshTokenRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		logErrorAndAbort(ctx, "refresh_token body err: %s", err)
		return
	}

	userId, err := crypto.ValidateTokenAndReturnUserId(req.RefreshToken)
	if err != nil {
		logger.Warn("validate refresh_token failed: %s", err);
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	clientInfoRaw, _ := ctx.Get(utils.CLIENT_INFO_CONTEXT_CONST)
	clientInfo := clientInfoRaw.(utils.ClientInfo)

	// check with temporary storage
	if _, err := h.tokenData.FindByToken(req.RefreshToken, clientInfo.Id); err != nil {
		logger.Warn("check refresh_token failed: %s", err);
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	authToken, refreshToken, expAt, tokenErr := generateTokens(userId)
	if tokenErr != nil {
		logErrorAndAbort(ctx, "refresh_token generate token err: %s", tokenErr)
		return
	}

	// revoke previous and save refresh_token in temporary storage (with TTL)
	newToken := data.Token{
		Token:          refreshToken,
		UserId:         userId,
		ExpirationTime: expAt,
		ClientInfo:     clientInfo.Name,
		ClientId:       clientInfo.Id,
	}
	if err = h.tokenData.Add(newToken); err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	h.tokenData.RemoveByToken(req.RefreshToken, clientInfo.Id)

	ctx.JSON(http.StatusOK, RefreshTokenResponse{
		authToken,
		refreshToken,
	})
}

func (h *httpHandler) AuthCreateAccount(ctx *gin.Context) {
	var req CreateAccountRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		logErrorAndAbort(ctx, "create_account body err: %s", err)
		return
	}

	// TODO add request validations

	username := strings.ToLower(req.Username)
	creationTime := time.Now().Unix()

	passwordHash := crypto.GeneratePasswordHash(req.Password, username, int(creationTime))
	if err := h.userData.Add(req.Username, passwordHash, creationTime); err != nil {
		logger.Error("add_user failed, err: %s", err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusOK)
}

// generateTokens return authToken, refreshToken, expirationTimeOfRefreshToken and error
func generateTokens(userId string) (string, string, int64, error) {
	authToken, _, err := crypto.GenereteToken(userId, 5) // 5 min
	if err != nil {
		return "", "", -1, err
	}
	refreshToken, exp, err := crypto.GenereteToken(userId, 7*60*24) // 7 days

	return authToken, refreshToken, exp, err
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
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenResponse struct {
	AuthToken    string `json:"auth_token"`
	RefreshToken string `json:"refresh_token"`
}

type CreateAccountRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
