package main

import (
	"bufio"
	"fmt"
	"os"
)

type User struct {
	Name  string
	Email string
	Age   int
	Role  string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("./users.txt") // Byte representasi data biner
	check(err)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line[0])
		break
	}

	// fmt.Println(string(user))

}
