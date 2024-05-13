package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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
