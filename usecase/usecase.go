package usecase

import "language-dictionary/model"

type languageUseCase struct{}

func NewLanguageUseCase() LanguageUseCase {
	return &languageUseCase{}
}

type LanguageUseCase interface {
	GetProgrammingLanguage() model.ProgrammingLanguage
}

func (uc *languageUseCase) GetProgrammingLanguage() model.ProgrammingLanguage {
	return model.ProgrammingLanguage{
		Language:       "C",
		Appeared:       1972,
		Created:        []string{"Dennis Ritchie"},
		Functional:     true,
		ObjectOriented: false,
		Relation: model.Relation{
			InfluencedBy: []string{"B", "ALGOL 68", "Assembly", "FORTRAN"},
			Influences:   []string{"C++", "Objective-C", "C#", "Java", "JavaScript", "PHP", "Go"},
		},
	}
}
