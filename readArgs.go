package main

import (
	"flag"
)

func readArgs() {
	quoteLengthArg := flag.Int("length", 10, "Enter the no. of words")
	stopOnErrorArg := flag.Bool("stop-on-error", false, "Set wheather to stop on error or not.")

	flag.Parse()

	quoteLength = *quoteLengthArg
	stopOnError = *stopOnErrorArg
}
