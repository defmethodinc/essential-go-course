package main

import (
	"fmt"
	"strings"
)

func GetLine(start int) string {
	line := make([]string, 10, 10)
	for i := 0; i < 10; i++ {
		line[i] = fmt.Sprintf("%d", start)
		start++
	}
	return strings.Join(line, " ")
}

func main() {
	fmt.Println(GetLine(0))
	fmt.Println(GetLine(10))
	fmt.Println(GetLine(20))
}
