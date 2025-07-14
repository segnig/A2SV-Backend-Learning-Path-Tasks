package main

import (
	"strings"
	"unicode"
)

func wordFrequency(input string) map[string]uint {
	wordCounter := make(map[string]uint)

	var builder strings.Builder

	for _, char := range input {
		if unicode.IsLetter(char) || unicode.IsNumber(char) || unicode.IsSpace(char) {
			builder.WriteRune(unicode.ToLower(char))
		}
	}

	words := strings.Fields(builder.String())

	for _, word := range words {
		wordCounter[word]++
	}

	return wordCounter
}
