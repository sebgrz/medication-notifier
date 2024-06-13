package main

import (
	"medication-notifier/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	apiAuth := router.Group("/api/auth")
	{
		apiAuth.POST("login", handler.AuthLogin)
	}
	router.Run(":8080")
}
