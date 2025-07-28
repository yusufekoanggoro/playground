package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type Printer struct{}

func (*Printer) Print(msg string) {
	fmt.Println(msg)
}

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

	// input bandung:yusuf,jakarta:eko,jawa:anggoro,jawa:udin
	input := "bandung: yusuf,jakarta:eko,jawa:anggoro,jawa:udin,jakarta:awan,jakarta:syamsul"

	userMap := map[string]string{}

	split := strings.Split(input, ",")

	for _, data := range split {
		tmp := strings.Split(data, ":")
		name := strings.TrimSpace(tmp[1])
		if userMap[tmp[0]] != "" {
			userMap[tmp[0]] += "," + name
		} else {
			userMap[tmp[0]] += name
		}
	}

	for key, data := range userMap {
		fmt.Print(key, ":", len(strings.Split(data, ",")))
		fmt.Println("")
	}

	for key, data := range userMap {
		fmt.Print(key, ":", data)
		fmt.Println("")
	}

	slc := []int{1, 5, 7, 10, 2, 3, 100, 300, 90}

	largest := slc[0]

	for i := range slc {
		if slc[i] > largest {
			largest = slc[i]
		}
	}

	fmt.Println("largest: ", largest)

	var wg sync.WaitGroup

	total := 2

	for i := 0; i < total; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Printf("%d. wait in 3s\n", i)
			time.Sleep(3 * time.Second)
		}()
	}

	wg.Wait()

	ch := make(chan string)

	go func() {
		ch <- "asd1"
		time.Sleep(2 * time.Second)
		close(ch)
	}()

	// go func() {
	// 	ch <- "asd2"
	// 	time.Sleep(2 * time.Second)
	// 	close(ch)
	// }()

	for val := range ch {
		fmt.Println(val)
	}

	// fmt.Println("ch: ", <-ch)

	printer := Printer{}
	printer.Print("Yusuf Eko Anggoro")

	count := 3
	for i := 0; i < 3; i++ {
		// fmt.Println(count)
		for j := 0; j < count; j++ {
			fmt.Print("*")
		}
		fmt.Println()
		count--
	}
}
