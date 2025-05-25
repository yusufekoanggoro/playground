package main

import "fmt"

func main() {
	n := 5
	x := 6

	deret := []int{n}

	for i := 0; i < x; i++ {
		deret = append(deret, deret[i]*2)
	}

	fmt.Println(deret)
}
