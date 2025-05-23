package main

import "fmt"

func selectionSort(arr [4]int, sort string) {
	if sort == "asc" {
		var min int
		var temp int

		for i := 0; i < len(arr); i++ {
			min = i
			for j := i + 1; j < len(arr); j++ {
				if arr[j] < arr[min] { // ascending
					min = j
				}
			}
			temp = arr[i]
			arr[i] = arr[min]
			arr[min] = temp
		}
	}

	if sort == "desc" {
		var max int
		var temp int

		for i := 0; i < len(arr); i++ {
			max = i
			for j := i + 1; j < len(arr); j++ {
				if arr[j] > arr[max] { // descending
					max = j
				}
			}
			temp = arr[i]
			arr[i] = arr[max]
			arr[max] = temp
		}
	}

	fmt.Println(arr)
}

func main() {
	array := [4]int{3, 5, 9, 4}
	selectionSort(array, "asc")
	selectionSort(array, "desc")
}
