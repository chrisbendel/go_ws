package utils

import (
	"testing"
)

func TestCheckForLetter(t *testing.T) {
	word := "TESTING"
	guessed := []rune{'T', 'E'}
	newWord := ReplaceLetters(word, guessed)

	if newWord != "TE?T???" {
		t.Errorf("String %s should be equal to %s after replacing letters.", word, newWord)
	}
}
