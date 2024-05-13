package usecase

import "language-dictionary/model"

type languageUseCase struct{}

func NewLanguageUseCase() LanguageUseCase {
	return &languageUseCase{}
}

type LanguageUseCase interface {
	GetProgrammingLanguage() model.ProgrammingLanguage
	IsPalindrome(text string) bool
}
