package main

import (
	"fmt"
)

// Fungsi utama
func main() {
	var klubA, klubB string
	var skorA, skorB int
	var hasil []string

	fmt.Print("Masukkan nama Klub A: ")
	fmt.Scanln(&klubA)
	fmt.Print("Masukkan nama Klub B: ")
	fmt.Scanln(&klubB)

	for pertandingan := 1; ; pertandingan++ {
		fmt.Printf("Pertandingan %d: ", pertandingan)
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			fmt.Println("Skor negatif terdeteksi, program dihentikan.")
			break
		}

		if skorA > skorB {
			hasil = append(hasil, klubA)
			fmt.Printf("Hasil %d: %s menang\n", pertandingan, klubA)
		} else if skorB > skorA {
			hasil = append(hasil, klubB)
			fmt.Printf("Hasil %d: %s menang\n", pertandingan, klubB)
		} else {
			hasil = append(hasil, "Draw")
			fmt.Printf("Hasil %d: Draw\n", pertandingan)
		}
	}

	fmt.Println("\nDaftar klub yang memenangkan pertandingan:")
	for i, pemenang := range hasil {
		if pemenang != "Draw" {
			fmt.Printf("Pertandingan %d: %s\n", i+1, pemenang)
		}
	}
}
