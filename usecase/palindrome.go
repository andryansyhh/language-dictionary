package usecase

func (uc *languageUseCase) IsPalindrome(text string) bool {
	runes := []rune(text)
	for i := 0; i < len(runes)/2; i++ {
		if runes[i] != runes[len(runes)-i-1] {
			return false
		}
	}
	return true
}
