package main

import "fmt"

func main() {
	data := []int{9, 2, 3, 10, 23}

	for i := 0; i < len(data); i++ {
		minIndex := i // 1
		for j := i + 1; j < len(data); j++ {
			if data[j] < data[minIndex] {
				minIndex = j
			}
		}
		temp := data[i]
		data[i] = data[minIndex]
		data[minIndex] = temp
	}

	fmt.Println(data)

	slice := make([]int, 5)
	fmt.Println("length of slice: ", len(slice))
}
