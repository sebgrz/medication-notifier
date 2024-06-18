package middleware

import (
	"medication-notifier/crypto"
	"medication-notifier/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const BearerPrefix = "Bearer "

func JwtAuthMiddleware() gin.HandlerFunc {
	return func (ctx *gin.Context) {
		// check auth_token
		authHeader := ctx.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, BearerPrefix) {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		
		token := strings.TrimPrefix(authHeader, BearerPrefix)
		userId, err := crypto.ValidateTokenAndReturnUserId(token)
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.Set(utils.USER_ID_CONST, userId)
	}
}
