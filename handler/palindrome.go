package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *LanguageHandler) CheckPalindrome(c *gin.Context) {
	text := c.Query("text")
	isPalindrome := h.useCase.IsPalindrome(text)
	if isPalindrome {
		c.String(http.StatusOK, "is Palindrome")
	} else {
		c.String(http.StatusBadRequest, "Not palindrome")
	}
}
