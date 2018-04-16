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
	}
	n := rand.Int() % len(words)
	return strings.ToUpper(words[n])
}
