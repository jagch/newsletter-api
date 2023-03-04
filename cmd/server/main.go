package main

import (
	"git.mcontigo.com/safeplay/newsletter-api/pkg/cors"
	"git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter/handler"
	"github.com/gin-gonic/gin"
)

// @contact.name                Grupo MContigo
// @title                       Newsletter API
// @version                     1.0
// @description                 Newsletter API
func main() {
	r := gin.Default()

	r.Use(cors.AllowCORS())

	newsletterHandler := handler.Build()

	libGroup := r.Group("/newsletter")
	libGroup.GET("subscriptions", newsletterHandler.Get)
	libGroup.POST("subscriptions", newsletterHandler.Post)

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
