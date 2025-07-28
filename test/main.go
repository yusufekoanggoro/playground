package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := "50"

	split := strings.Split(input, "")

	satuan := make(map[int]string)
	satuan[1] = "Satu"

	puluhan := make(map[int]string)
	puluhan[5] = "Lima Puluh"
	puluhan[5] = "Lima Puluh"
	puluhan[5] = "Lima Puluh"
	puluhan[5] = "Lima Puluh"

	ratusan := make(map[int]string)
	ratusan[2] = "Dua Ratus"

	for i := 0; i < len(split); i++ {
		// puluhan
		if len(input) == 2 {
			v, _ := strconv.Atoi(split[i])
			fmt.Print(puluhan[v])
		}
		break
		// // ratusan
		// if len(input) == 3 {

		// }
		// // seribu
		// if len(input) == 3 {

		// }
	}
}
