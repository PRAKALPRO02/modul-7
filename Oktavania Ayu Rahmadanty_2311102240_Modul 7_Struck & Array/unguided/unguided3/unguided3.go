package main

import (
	"fmt"
)

func main() {
	var klubA, klubB string
	var skorA, skorB int
	pemenang := []string{}

	// Meminta input nama klub
	fmt.Print("Klub A: ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B: ")
	fmt.Scan(&klubB)

	// Meminta input skor pertandingan
	for i := 1; ; i++ {
		fmt.Printf("Pertandingan %d: ", i)
		fmt.Scan(&skorA, &skorB)

		// Menghentikan input jika ada skor negatif
		if skorA < 0 || skorB < 0 {
			fmt.Println("Pertandingan selesai")
			break
		}

		// Menentukan pemenang
		if skorA > skorB {
			pemenang = append(pemenang, klubA)
			fmt.Printf("Hasil %d: %s\n", i, klubA)
		} else if skorA < skorB {
			pemenang = append(pemenang, klubB)
			fmt.Printf("Hasil %d: %s\n", i, klubB)
		} else {
			fmt.Printf("Hasil %d: Draw\n", i)
		}
	}

	// Menampilkan daftar pemenang pertandingan
	fmt.Println("\nDaftar klub yang memenangkan pertandingan:")
	for _, klub := range pemenang {
		fmt.Println(klub)
	}
}
