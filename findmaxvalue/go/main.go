package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	// with array
	arr := [4]int{1, 89, 100, 82}

	largest := arr[0]

	for i := 0; i < len(arr); i++ {
		if arr[i] > largest {
			largest = arr[i]
		}
	}

	fmt.Println("with array: ", largest)

	// with slice
	slc := []int{1, 88, 999, 20}

	largest = slc[0]

	for i := 0; i < len(arr); i++ {
		if slc[i] > largest {
			largest = slc[i]
		}
	}

	fmt.Println("with slice: ", largest)

	var users []User

	users = append(users, User{Name: "Udin 1", Age: 50}) // udin 1
	users = append(users, User{Name: "Udin 2", Age: 39}) // udin 2
	users = append(users, User{Name: "Udin 3", Age: 20}) // udin 3
}
