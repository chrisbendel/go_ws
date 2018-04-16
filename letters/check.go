package letters

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

func IsCorrect(correctAnswer string, userAnswer string) bool {
	return correctAnswer == userAnswer
}
