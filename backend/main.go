package main

import (
	"medication-notifier/handler"
	"medication-notifier/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	// Auth endpoints
	apiAuth := router.Group("/api/auth")
	{
		apiAuth.POST("login", handler.AuthLogin)
		apiAuth.POST("refresh_token", handler.AuthRefreshToken)
	}

	// Application endpoints
	apiApp := router.Group("/api")
	apiApp.Use(middleware.JwtAuthMiddleware())
	{
		apiApp.GET("test", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "OK")
		})
	}

	router.Run(":8080")
}
