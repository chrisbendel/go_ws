package letters

import (
	"log"
	"regexp"
)

func ReplaceLetters(answer string, guessed []rune) string {
	var updatedAnswer = ""
	for _, c := range answer {
		if c == 32 {
			updatedAnswer += " "
		} else {
			updatedAnswer += "?"
		}
	}
	questionMarks := []rune(updatedAnswer)

	for answerIndex, character := range answer {
		for _, guess := range guessed {
			if character == guess {
				questionMarks[answerIndex] = character
			}
		}
	}

	return string(questionMarks)
}

func UserGuessToRune(userGuess string) rune {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")

	if err != nil {
		log.Fatal(err)
	}

	processedString := reg.ReplaceAllString(userGuess, "")
	guessedRune := rune('a')

	if processedString != "" {
		guessedRune = rune(processedString[0])
	}

	return guessedRune
}

func IsCorrect(correctAnswer string, userAnswer string) bool {
	return correctAnswer == userAnswer
}
