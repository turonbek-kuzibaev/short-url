package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/turonbek-kuzibaev/go-url-shortener/handler"
	"github.com/turonbek-kuzibaev/go-url-shortener/store"
)

func main() {
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hey Go URL Shortener !",
		})
	})

	r.POST("/create-short-url", func(ctx *gin.Context) {
		handler.CreateShortUrl(ctx)
	})

	r.GET("/:shortUrl", func(ctx *gin.Context) {
		handler.HandleShortUrlRedirect(ctx)
	})

	store.InitializeStore()

	err := r.Run(":8000")

	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}
