package main

import "time"

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
	typedLength int
	stopOnError bool = false
)
