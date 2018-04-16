package letters

import (
	"log"
	"testing"
)

func TestCheckForLetter(t *testing.T) {
	word := "testing"
	guessed := []rune{'t', 'e'}
	newWord := ReplaceLetters(word, guessed)

	log.Println(newWord)
	if newWord != "te?t???" {
		t.Errorf("String %s should be equal to %s after replacing letters.", word, newWord)
	}
}
