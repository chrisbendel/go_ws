package words

import (
	"math/rand"
	"strings"
	"time"
)

func GetRandomWord() string {
	rand.Seed(time.Now().Unix())
	words := []string{
		"TESTING the cool thing",
		"PHRASE with space",
		"GOLANG fun",
		"TWO cool dogs",
		"this is fun times",
		"hangman is a great game",
		"bud light is not so bad",
		"caterpillar",
	}
	n := rand.Int() % len(words)
	return strings.ToUpper(words[n])
}
