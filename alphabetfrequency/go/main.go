package main

import (
	"fmt"
	"strings"
)

func main() {
	word := "backend developer"

	counter := make(map[string]int)

	split := strings.Split(strings.ReplaceAll(word, " ", ""), "")

	for i := 0; i < len(split); i++ {
		counter[split[i]]++
	}

	fmt.Println(counter)
}
