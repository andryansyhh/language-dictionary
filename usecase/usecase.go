package usecase

import "language-dictionary/model"

type languageUseCase struct {
	languages []model.ProgrammingLanguage
}

func NewLanguageUseCase() LanguageUseCase {
	return &languageUseCase{}
}

type LanguageUseCase interface {
	// GetProgrammingLanguage() model.ProgrammingLanguage
	IsPalindrome(text string) bool
	AddLanguage(language model.ProgrammingLanguage) (model.ProgrammingLanguage, error)
	GetLanguageByID(id int) (model.ProgrammingLanguage, error)
	GetAllLanguages() []model.ProgrammingLanguage
	UpdateLanguage(id int, language model.ProgrammingLanguage) (model.ProgrammingLanguage, error)
	DeleteLanguage(id int) error
}
