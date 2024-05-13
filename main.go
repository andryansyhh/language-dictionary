package main

import (
	"language-dictionary/handler"
	"language-dictionary/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Initialize use case
	usecase := usecase.NewLanguageUseCase()

	// Initialize handler with use case
	handler := handler.NewLanguageHandler(usecase)

	r.GET("/", handler.HelloHandler)
	r.GET("/language", handler.GetLanguage)
	r.NoRoute(handler.MethodNotAllowedHandler)

	r.Run(":7777")
}
