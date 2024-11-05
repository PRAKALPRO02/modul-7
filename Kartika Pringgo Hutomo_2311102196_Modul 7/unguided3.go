package main

import "fmt"

func main() {
	var klubA, klubB string
	var skorA, skorB int
	var hasil []string

	fmt.Print("Nama Klub A: ")
	fmt.Scanln(&klubA)
	fmt.Print("Nama Klub B: ")
	fmt.Scanln(&klubB)

	pertandingan := 1
	for {
		fmt.Printf("Skor Pertandingan %d (Klub A Klub B): ", pertandingan)
		fmt.Scanln(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			hasil = append(hasil, klubA)
		} else if skorA < skorB {
			hasil = append(hasil, klubB)
		} else {
			hasil = append(hasil, "Draw")
		}

		pertandingan++
	}

	fmt.Println("Pertandingan selesai")
	for i, p := range hasil {
		fmt.Printf("Hasil Pertandingan %d: %s\n", i+1, p)
	}
}
