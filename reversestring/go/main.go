package main

import (
	"fmt"
	"strings"
)

func reverseString(str string) {
	split := strings.Split(str, "")
	for i := len(split) - 1; i >= 0; i-- {
		fmt.Print(split[i])
	}
}

func main() {
	reverseString("Yusuf Eko Anggoro")
}
