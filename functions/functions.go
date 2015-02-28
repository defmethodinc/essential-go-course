package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Double(5))

	first, _ := ParseName("Joe Leo")
	fmt.Println(first)

	greet := func(name string) {
		fmt.Println("Hello", name)
	}
	greet(first)

}

func Double(n int) int {
	return n + n
}

func ParseName(name string) (first, last string) {
	parsed := strings.Split(name, " ")
	return parsed[0], parsed[1]
}
