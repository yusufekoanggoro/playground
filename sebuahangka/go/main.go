package main

import (
	"fmt"
	"strings"
)

func main() {
	// Ada sebuah angka seperti ini :
	// 1.225.441

	// Berikan outputnya berupa :
	// 1000000
	// 200000
	// 20000
	// 5000
	// 400
	// 40
	// 1
	input := "225.441"

	cleaned := strings.ReplaceAll(input, ".", "")

	split := strings.Split(cleaned, "")

	n := len(split) - 1
	for i := 0; i < len(split); i++ {
		fmt.Print(split[i])
		for j := n; j > 0; j-- {
			fmt.Print("0")
		}
		n--
		fmt.Println()
	}
}
