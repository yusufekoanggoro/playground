package main

import (
	"fmt"
	"strings"
)

func main() {
	word := "halo halo bandung ibu kota periangan"

	split := strings.Split(word, " ")

	wordCount := make(map[string]int)

	for i := 0; i < len(split); i++ {
		wordCount[split[i]]++
	}

	fmt.Println(wordCount)
}
