package handler

import (
	"language-dictionary/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LanguageHandler struct {
	useCase usecase.LanguageUseCase
}

func NewLanguageHandler(useCase usecase.LanguageUseCase) *LanguageHandler {
	return &LanguageHandler{useCase: useCase}
}

func (h *LanguageHandler) GetLanguage(c *gin.Context) {
	language := h.useCase.GetProgrammingLanguage()
	c.JSON(http.StatusOK, language)
}

func (h *LanguageHandler) HelloHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello Go developers")
}

func (h *LanguageHandler) MethodNotAllowedHandler(c *gin.Context) {
	c.String(http.StatusMethodNotAllowed, "Method not allowed")
}
