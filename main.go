package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"golang.org/x/term"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
)

var (
	currPos     int = 0
	startTime   time.Time
	currColor   string = Reset
	quoteLength int
)

func main() {
	if len(os.Args) >= 2 {
		length, err := strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatal("hahaha")
		}
		quoteLength = length
	} else {
		quoteLength = 10
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	charChan := make(chan byte)

	go func() {
		buf := make([]byte, 1)
		for {
			_, err := os.Stdin.Read(buf)
			if err != nil {
				return
			}
			charChan <- buf[0]
		}
	}()

	quote := getQuote(quoteLength)

	for {
		select {
		case char := <-charChan:

			width, height, err := term.GetSize(int(os.Stdout.Fd()))
			if err != nil {
				log.Fatalf("Could not get terminal size: %v", err)
			}
			if char == 3 {
				fmt.Print("\r\nManual Ctrl+C detected. Exiting...\r\n")
				return
			}
			if currPos >= len(quote)-1 {
				endTime := time.Since(startTime)
				wpm := float64(len(quote)/5) / endTime.Minutes()

				fmt.Print("\033[2J")
				fmt.Printf("\033[%d;%dH", height/2, (width/2)-7)
				fmt.Printf("Your WPM is %0.f , ", wpm)
				fmt.Print("Press e to start again!")
				if string(char) == "e" {
					quote = getQuote(quoteLength)
					currPos = 0
				}
				continue
			}
			if currPos == 0 {
				startTime = time.Now()
			}

			if char == quote[currPos] {
				currPos++
				currColor = Green
			}
			if char != quote[currPos] {
				currColor = Red
			}

			typedQuote := ""
			if currPos == 0 {
				typedQuote = ""
			} else {
				typedQuote = quote[0:currPos]
			}

			remainingQuote := quote[currPos+1:]
			currQuote := quote[currPos]

			fmt.Print("\033[2J")
			fmt.Printf("\033[%d;%dH", height/2, (width/2)-len(quote)/2)

			fmt.Print(Green + typedQuote + Reset + currColor + string(currQuote) + Reset + remainingQuote)

			fmt.Printf("\033[%d;%dH", height/2, (width/2)-len(quote)/2+currPos)

		case <-sigs:
			// This catches the standard OS signal
			fmt.Print("\r\nInterrupt signal received. Exiting...\r\n")
			return
		}
	}
}
