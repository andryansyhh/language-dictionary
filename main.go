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

	r.NoRoute(handler.MethodNotAllowedHandler)

	// modified endpoint
	// r.GET("/", handler.HelloHandler)
	// r.GET("/language", handler.GetLanguage)

	// palindrome
	r.GET("/palindrome", handler.CheckPalindrome)

	// language-dictionary
	r.POST("/language", handler.AddLanguage)
	r.GET("/language/:id", handler.GetLanguageByID)
	r.GET("/languages", handler.GetAllLanguages)
	r.PATCH("/language/:id", handler.UpdateLanguage)
	r.DELETE("/language/:id", handler.DeleteLanguage)

	r.Run(":7777")
}
