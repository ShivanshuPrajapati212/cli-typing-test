package main

import "time"

const (
	Reset  = "\033[38;2;86;95;135m"
	Red    = "\033[38;2;247;118;142m"
	Green  = "\033[38;2;192;202;245m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
)

var (
	currPos     int = 0
	startTime   time.Time
	currColor   string = Reset
	quoteLength int
	typedLength int
	stopOnError bool = false
	typedQuote  string
)
