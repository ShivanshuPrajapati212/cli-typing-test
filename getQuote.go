package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"slices"
)

func getQuote(count int) string {
	file, err := os.Open("wordlist.txt")
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer file.Close() // Ensure the file closes when we are done

	var quote string

	randomNumbers := make([]int, count)

	for i := range count {
		// Generates numbers between 0 and 99
		randomNumbers[i] = rand.IntN(200)
	}

	scanner := bufio.NewScanner(file)
	currLine := 1
	for scanner.Scan() {
		if slices.Contains(randomNumbers, currLine) {
			quote = quote + scanner.Text() + " "
		}
		currLine++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error during scan:", err)
	}

	return quote
}
