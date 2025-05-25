package main

import "fmt"

// TODO: STRUCT DATA
type Candidate struct {
	Name     string
	Location string
}

var registeredNames []string

func isNameRegistered(name string) bool {
	for _, registeredName := range registeredNames {
		if registeredName == name {
			return true
		}
	}
	return false
}

// TODO: ADD REGISTER LOGIC
func register(candidate Candidate) {
	if candidate.Location == "" {
		fmt.Println("Gagal mendaftar, Pesan: Data Peserta tidak lengkap.")
		return
	}

	if isNameRegistered(candidate.Name) {
		fmt.Println("Gagal mendaftar, Pesan: Nama sudah terdaftar.")
		return
	}

	registeredNames = append(registeredNames, candidate.Name)

	fmt.Println("Berhasil mendaftar!")
}

func executeProgram(candidates []Candidate) {
	if len(candidates) > 5 {
		fmt.Print("Gagal mendaftar, Pesan: Kuota telah terpenuhi.")
	}

	// var isExist bool = false
	for i := 0; i < len(candidates); i++ {
		fmt.Print("Peserta ", i+1, ": ")
		register(candidates[i])
	}
}

func main() {
	// TODO: ADD YOUR CANDIDATES DATA
	candidates := []Candidate{}

	candidates = append(candidates, Candidate{Name: "Yusuf", Location: "Jakarta Timur"})
	candidates = append(candidates, Candidate{Name: "Yusuf", Location: "Jakarta Timur"})
	candidates = append(candidates, Candidate{Name: "Anggoro", Location: "Jakarta Timur"})

	// START PROGRAM
	executeProgram(candidates)
}
