package main

import (
	"medication-notifier/data/db"
	"medication-notifier/handler"
	"medication-notifier/middleware"

	"github.com/gin-gonic/gin"
)

func main() {

	userDataService := db.NewDummyUsersDataService()
	tokenDataService := db.NewDbTokenDataService("localhost:6379", "")
	medicationDataService := db.NewDummyMedicationDataService()
	handler := handler.New(&userDataService, &tokenDataService, &medicationDataService)

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
		apiApp.GET("list", handler.ListMedications)
		apiApp.POST("add", handler.AddMedication)
		apiApp.DELETE("remove/:id", handler.RemoveMedication)
		apiApp.PUT("replace", handler.ReplaceMedication)
	}

	router.Run(":8080")
}
