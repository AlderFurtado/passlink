package utils

import "math/rand"

func RandomWord(length int) string {
	letters := "abcdefghijklmnopqrstuvwxyz"
	word := make([]byte, length)
	for i := range word {
		word[i] = letters[rand.Intn(len(letters))]
	}
	return string(word)
}
