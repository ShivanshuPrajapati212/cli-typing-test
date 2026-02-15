package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/term"
)

var quote string = "hello bro you are nice man, and i like you to be my friend"

func main() {
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		log.Fatalf("Could not get terminal size: %v", err)
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	fmt.Print("\033[2J")

	fmt.Printf("\033[%d;%dH", height/2, (width/2)-len(quote)/2)

	fmt.Print(quote)

	sig := <-sigs
	fmt.Print(sig)
}
