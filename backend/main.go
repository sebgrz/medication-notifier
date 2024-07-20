package main

import (
	"context"
	"fmt"
	"medication-notifier/data/db"
	"medication-notifier/handler"
	"medication-notifier/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	sqlAddress := "postgres://medication:medication@localhost:5432/medication_db?sslmode=disable"
	db.RunMigration(sqlAddress)

	conn, err := pgxpool.New(context.Background(), sqlAddress)
	if err != nil {
		panic(fmt.Sprintf("psql connection failed: %s", err))
	}

	userDataService := db.NewDbUsersDataService(conn)
	medicationDataService := db.NewDbMedicationDataService(conn)
	tokenDataService := db.NewDbTokenDataService("localhost:6379", "")
	handler := handler.New(&userDataService, &tokenDataService, &medicationDataService)

	router := gin.New()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "DELETE", "PUT"},
		AllowHeaders:     []string{"Origin", "X-Client-Id", "User-Agent", "Authorization"},
		AllowCredentials: true,
	}))
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
