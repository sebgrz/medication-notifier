package main

import (
	"medication-notifier/data/db"
	"medication-notifier/handler"
	"medication-notifier/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	userDataService := db.NewDummyUsersDataService()
	tokenDataService := db.NewDummyTokenDataService()
	handler := handler.New(&userDataService, &tokenDataService)

	router := gin.New()
	router.Use(middleware.ClientInfoMiddleware())

	// Auth endpoints
	apiAuth := router.Group("/api/auth")
	{
		apiAuth.POST("login", handler.AuthLogin)
		apiAuth.POST("refresh_token", handler.AuthRefreshToken)
		apiAuth.POST("register", handler.AuthCreateAccount)
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
