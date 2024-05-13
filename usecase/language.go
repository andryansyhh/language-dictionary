package usecase

import (
	"errors"
	"language-dictionary/model"
)

// func (uc *languageUseCase) GetProgrammingLanguage() model.ProgrammingLanguage {
// 	return model.ProgrammingLanguage{
// 		Language:       "C",
// 		Appeared:       1972,
// 		Created:        []string{"Dennis Ritchie"},
// 		Functional:     true,
// 		ObjectOriented: false,
// 		Relation: model.Relation{
// 			InfluencedBy: []string{"B", "ALGOL 68", "Assembly", "FORTRAN"},
// 			Influences:   []string{"C++", "Objective-C", "C#", "Java", "JavaScript", "PHP", "Go"},
// 		},
// 	}
// }

func (uc *languageUseCase) AddLanguage(language model.ProgrammingLanguage) (model.ProgrammingLanguage, error) {
	language.Id = len(uc.languages)
	uc.languages = append(uc.languages, language)
	return language, nil
}

func (uc *languageUseCase) GetLanguageByID(id int) (model.ProgrammingLanguage, error) {
	if id < 0 || id >= len(uc.languages) {
		return model.ProgrammingLanguage{}, errors.New("invalid language id")
	}
	return uc.languages[id], nil
}

func (uc *languageUseCase) GetAllLanguages() []model.ProgrammingLanguage {
	return uc.languages
}

func (uc *languageUseCase) UpdateLanguage(id int, language model.ProgrammingLanguage) (model.ProgrammingLanguage, error) {
	if id < 0 || id >= len(uc.languages) {
		return model.ProgrammingLanguage{}, errors.New("invalid language id")
	}

	uc.languages[id] = language

	uc.languages[id].Id = id

	return uc.languages[id], nil
}

func (uc *languageUseCase) DeleteLanguage(id int) error {
	if id < 0 || id >= len(uc.languages) {
		return errors.New("invalid language ID")
	}

	uc.languages = append(uc.languages[:id], uc.languages[id+1:]...)

	for i := id; i < len(uc.languages); i++ {
		uc.languages[i].Id = i
	}

	return nil
}
