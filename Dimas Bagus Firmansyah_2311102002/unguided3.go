package main

import (
	"fmt"
)

func main() {
	// Memasukkan nama kedua klub
	var klubA, klubB string
	fmt.Print("Klub A: ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B: ")
	fmt.Scan(&klubB)

	// Array untuk menyimpan hasil klub yang menang
	var pemenang []string

	// Input skor pertandingan
	var skorA, skorB int
	pertandingan := 1

	for {
		fmt.Printf("Pertandingan %d: ", pertandingan)
		fmt.Scan(&skorA, &skorB)

		// Hentikan input jika skor tidak valid (negatif)
		if skorA < 0 || skorB < 0 {
			break
		}

		// Menentukan pemenang berdasarkan skor
		if skorA > skorB {
			pemenang = append(pemenang, klubA)
		} else if skorA < skorB {
			pemenang = append(pemenang, klubB)
		} else {
			pemenang = append(pemenang, "Draw")
		}

		pertandingan++
	}

	// Menampilkan hasil setiap pertandingan
	for i, hasil := range pemenang {
		fmt.Printf("Hasil %d: %s\n", i+1, hasil)
	}

	fmt.Println("Pertandingan selesai")
}
