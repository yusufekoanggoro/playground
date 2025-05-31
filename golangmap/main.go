package main

import (
	"fmt"
	"strings"
)

type User struct {
	Name string
	Age  int
}

func main() {
	userMap := []map[string]User{}

	userMap = append(userMap, map[string]User{"1": {Name: "Yusuf", Age: 1}})

	for _, data := range userMap {
		fmt.Println(data["1"].Name)
	}

	word := "halo halo bandung ibu kota periangan"

	wordCount := map[string]int{}

	for _, data := range strings.Split(word, " ") {
		wordCount[data]++
	}

	fmt.Println(wordCount)
}
