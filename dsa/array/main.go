package main

import "fmt"

var data []int

func main() {
	data = append(data, 2)
	data = append(data, 3)
	data = append(data, 4)
	data = append(data, 5)

	data = append(data[:2], data[2+1:]...)
	fmt.Println(data)
}
