package handler

import (
	"language-dictionary/usecase"
)

type LanguageHandler struct {
	useCase usecase.LanguageUseCase
}

func NewLanguageHandler(useCase usecase.LanguageUseCase) *LanguageHandler {
	return &LanguageHandler{useCase: useCase}
}
