package main

import (
	"log"
	"os"
	"strconv"
)

func readArgs() {
	if len(os.Args) >= 2 {
		length, err := strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatal("hahaha")
		}
		quoteLength = length
	} else {
		quoteLength = 10
	}
}
