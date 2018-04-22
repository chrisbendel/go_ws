package utils

import (
	"testing"
	"unicode"
)

func TestReplaceLetters(t *testing.T) {
	word := "TESTING"
	guessed := []rune{'T', 'E'}
	newWord := ReplaceLetters(word, guessed)

	if newWord != "TE?T???" {
		t.Errorf("String %s should be equal to %s after replacing letters.", word, newWord)
	}
}

func TestUserGuessToRune(t *testing.T) {
	guessString := "'c'"
	targetRune := 'c'
	guessRune := UserGuessToRune(guessString)

	if unicode.IsLetter(guessRune) && guessRune != targetRune {
		t.Errorf("Rune %c should be %c", guessRune, targetRune)
	}
}

func TestIsCorrect(t *testing.T) {
	userAnswer := "WORD"
	correctAnswer := "WORD"

	if !IsCorrect(correctAnswer, userAnswer) {
		t.Errorf("User's answer %s should be equal to %s", userAnswer, correctAnswer)
	}
}

func TestGetRandomWord(t *testing.T) {
	randomWord := GetRandomWord()
	switch randomWord {
	case
		"TESTING THE COOL THING",
		"PHRASE WITH SPACE",
		"GOLANG FUN",
		"TWO COOL DOGS",
		"THIS IS FUN TIMES",
		"HANGMAN IS A GREAT GAME",
		"BUD LIGHT IS NOT SO BAD",
		"CATERPILLAR":
		break
	default:
		t.Errorf("%s is not an available word", randomWord)
	}

}
