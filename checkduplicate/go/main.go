package main

import "fmt"

func checkDuplicate(data []string) bool {
	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			if data[i] == data[j] {
				return true
			}
		}
	}
	return false
}

func checkDuplicateWithMap(data []string) bool {
	var m = make(map[string]string)
	for i := 0; i < len(data); i++ {
		var _, isExist = m[data[i]]
		if !isExist {
			m[data[i]] = data[i]
		} else {
			return true
		}
	}
	return false
}

func main() {
	isDuplicate := checkDuplicateWithMap([]string{"empat", "satu", "dua", "tiga", "lima"})
	if isDuplicate {
		fmt.Println("data duplicated")
	} else {
		fmt.Println("not duplicated")
	}

	var a = []int{1, 3, 4, 5, 6, 7}

	myslice := []int{}

	for i := range a {
		if i != 2 {
			myslice = append(myslice, a[i])
		}
	}

	fmt.Println(myslice)
}
