package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "2500000"

	split := strings.Split(input, "")
	length := len(input)

	digits := length - 1

	for i := 0; i < length; i++ {
		if split[i] != "0" {
			fmt.Print(split[i])
			for j := 0; j < digits; j++ {
				fmt.Print("0")
			}
			fmt.Println("")
		}
		digits--
	}
}
