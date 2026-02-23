package main

import (
	"math/rand/v2"
	"strings"
)

func getQuote(count int) string {
	var quote string

	randomNumbers := make([]int, count)

	for i := range count {
		// Generates numbers between 0 and 99
		randomNumbers[i] = rand.IntN(200)
	}

	words := strings.Split(WORDLIST, "\n")

	for _, v := range randomNumbers {
		quote = quote + words[v] + " "
	}

	return quote
}
