package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		fmt.Println("waiting in 2 seconds")
		time.Sleep(2 * time.Second)
	}()

	go func() {
		defer wg.Done()
		fmt.Println("waiting in 2 seconds")
		time.Sleep(2 * time.Second)
	}()

	wg.Wait()

	name := make(chan string, 1)

	go func() {
		name <- "Yusuf Eko Anggoro"
	}()

	fmt.Println("My Name: ", <-name)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer func() {
		defer cancel()
		if r := recover(); r != nil {
			fmt.Println("failed: ", r)
		}
	}()

	data := make(chan string)
	go func() {
		defer close(data)
		time.Sleep(6 * time.Second)
		data <- "as"
	}()

	select {
	case <-data:
		fmt.Println("received data")
	case <-ctx.Done():
		panic("timeout")
	}

	fmt.Println("end")
}
