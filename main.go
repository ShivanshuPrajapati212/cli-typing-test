package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/term"
)

var (
	quote   string = "hello bro you are nice man, and i like you to be my friend"
	currPos int    = 0
)

func main() {
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		log.Fatalf("Could not get terminal size: %v", err)
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

	for {
		select {
		case char := <-charChan:
			if char == 3 {
				fmt.Print("\r\nManual Ctrl+C detected. Exiting...\r\n")
				return
			}
			if currPos >= len(quote)-1 {
				fmt.Print("You did it.")
				return
			}

			if char == quote[currPos] {
				currPos++
			}

			fmt.Print("\033[2J")

			fmt.Printf("\033[%d;%dH", height/2, (width/2)-len(quote)/2)

			fmt.Print(quote)

			fmt.Printf("\033[%d;%dH", height/2, (width/2)-len(quote)/2+currPos)

		case <-sigs:
			// This catches the standard OS signal
			fmt.Print("\r\nInterrupt signal received. Exiting...\r\n")
			return
		}
	}
}
