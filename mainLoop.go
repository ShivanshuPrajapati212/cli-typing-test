package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"unicode"

	"golang.org/x/term"
)

func mainLoop() {
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
			if string(char) == "\n" || string(char) == "\r" {
				quote = getQuote(quoteLength)
				currPos = 0
				typedLength = 0
				typedQuote = ""
			}
			if currPos >= len(quote)-1 {
				endTime := time.Since(startTime)
				wpm := float64(typedLength/5) / endTime.Minutes()

				resLength := len(fmt.Sprintf("Your WPM is %0.f\nPress ENTER to start again!", wpm))
				fmt.Print("\033[2J")
				fmt.Printf("\033[%d;%dH", height/2, (width/2)-resLength/2)
				fmt.Printf("Your WPM is %0.f, Press ENTER to start again!", wpm)

				if string(char) == "\n" || string(char) == "\r" {
					quote = getQuote(quoteLength)
					currPos = 0
					typedLength = 0
					typedQuote = ""
				}
				continue
			}
			if currPos == 0 {
				startTime = time.Now()
			}

			if char == quote[currPos] {
				typedQuote = typedQuote + Green + string(quote[currPos]) + Reset
				typedLength++
				currPos++
				currColor = Green
			} else {
				if !stopOnError && unicode.IsPrint(rune(char)) && !unicode.IsControl(rune(char)) {
					typedQuote = typedQuote + Red + string(quote[currPos]) + Reset
					currPos++
				}
				currColor = Red
			}

			remainingQuote := quote[currPos+1:]
			currQuote := quote[currPos]

			fmt.Print("\033[2J")
			fmt.Printf("\033[%d;%dH", height/2, (width/2)-len(quote)/2)

			fmt.Print(typedQuote + currColor + string(currQuote) + Reset + remainingQuote)

			fmt.Printf("\033[%d;%dH", height/2, (width/2)-len(quote)/2+currPos)

		case <-sigs:
			// This catches the standard OS signal
			fmt.Print("\r\nInterrupt signal received. Exiting...\r\n")
			return
		}
	}
}
