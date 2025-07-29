package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "1234"

	split := strings.Split(input, "")

	satuan := make(map[string]string)
	satuan["1"] = "Satu"
	satuan["2"] = "Dua"
	satuan["3"] = "Tiga"
	satuan["4"] = "Empat"
	satuan["5"] = "Lima"
	satuan["6"] = "Enam"
	satuan["7"] = "Tujuh"
	satuan["8"] = "Delapan"
	satuan["9"] = "Sembilan"

	puluhan := map[string]string{
		"1": "Sepuluh",
		"2": "Dua Puluh",
		"3": "Tiga Puluh",
		"4": "Empat Puluh",
		"5": "Lima Puluh",
		"6": "Enam Puluh",
		"7": "Tujuh Puluh",
		"8": "Delapan Puluh",
		"9": "Sembilan Puluh",
	}

	ratusan := map[string]string{
		"1": "Seratus",
		"2": "Dua Ratus",
		"3": "Tiga Ratus",
		"4": "Empat Ratus",
		"5": "Lima Ratus",
		"6": "Enam Ratus",
		"7": "Tujuh Ratus",
		"8": "Delapan Ratus",
		"9": "Sembilan Ratus",
	}

	ribuan := map[string]string{
		"1": "Seribu",
		"2": "Dua Ribu",
		"3": "Tiga Ribu",
		"4": "Empat Ribu",
		"5": "Lima Ribu",
		"6": "Enam Ribu",
		"7": "Tujuh Ribu",
		"8": "Delapan Ribu",
		"9": "Sembilan Ribu",
	}

	var merged []map[string]string
	merged = append(merged, satuan)
	merged = append(merged, puluhan)
	merged = append(merged, ratusan)
	merged = append(merged, ribuan)

	length := len(split)
	digits := length - 1

	for i := 0; i < length; i++ {
		if split[i] != "0" {
			fmt.Print(merged[digits][split[i]], " ")
		}
		digits--
	}
}
