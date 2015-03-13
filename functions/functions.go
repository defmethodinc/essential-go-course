package main

import (
	"bytes"
	"fmt"
	"log"
	"strings"
)

var buf bytes.Buffer
var LoggerName string = "MyLogger"
var logger = log.New(&buf, LoggerName + ": ", log.Lshortfile)

func main() {
	fmt.Println(Double(5))

	first, _ := ParseName("Joe Leo")
	fmt.Println(first)

	greet := func(name string) {
		fmt.Println("Hello", name)
	}
	greet(first)

	line := buildLogMessage("Hello, log file!")
	fmt.Print(line)
}

func buildLogMessage(msg string) (string) {
	logger.Print(msg)
	return buf.String()
}

func Double(n int) int {
	return n + n
}

func ParseName(name string) (first, last string) {
	parsed := strings.Split(name, " ")
	return parsed[0], parsed[1]
}
