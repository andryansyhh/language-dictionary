package handler

import (
	"language-dictionary/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// func (h *LanguageHandler) GetLanguage(c *gin.Context) {
// 	language := h.useCase.GetProgrammingLanguage()
// 	c.JSON(http.StatusOK, language)
// }

func (h *LanguageHandler) HelloHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello Go developers")
}

func (h *LanguageHandler) MethodNotAllowedHandler(c *gin.Context) {
	c.String(http.StatusMethodNotAllowed, "Method not allowed")
}

func (h *LanguageHandler) AddLanguage(c *gin.Context) {
	var newLanguage model.ProgrammingLanguage
	if err := c.BindJSON(&newLanguage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	language, err := h.useCase.AddLanguage(newLanguage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, language)
}

func (h *LanguageHandler) GetLanguageByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid language ID"})
		return
	}

	language, err := h.useCase.GetLanguageByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, language)
}

func (h *LanguageHandler) GetAllLanguages(c *gin.Context) {
	languages := h.useCase.GetAllLanguages()
	c.JSON(http.StatusOK, languages)
}

func (h *LanguageHandler) UpdateLanguage(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid language ID"})
		return
	}

	var updatedLanguage model.ProgrammingLanguage
	if err := c.BindJSON(&updatedLanguage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	language, err := h.useCase.UpdateLanguage(id, updatedLanguage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, language)
}

func (h *LanguageHandler) DeleteLanguage(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid language ID"})
		return
	}

	err = h.useCase.DeleteLanguage(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Language deleted"})
}
