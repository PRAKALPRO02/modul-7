package main

import (
	"fmt"
	"strings"
)

func main() {
	var klub1, klub2 string
	var skor1, skor2 int
	var pemenang []string

	// Meminta masukan nama klub yang bertanding
	fmt.Print("Masukkan nama Klub A: ")
	fmt.Scan(&klub1)
	fmt.Print("Masukkan nama Klub B: ")
	fmt.Scan(&klub2)

	// Mulai mencatat hasil pertandingan
	for i := 1; ; i++ {
		fmt.Printf("Pertandingan %d (skor %s vs %s): ", i, klub1, klub2)
		n, err := fmt.Scan(&skor1, &skor2)
		
		// Cek jika input bukan angka atau skor negatif (menghentikan program)
		if err != nil || n != 2 || skor1 < 0 || skor2 < 0 {
			fmt.Println("Input skor tidak valid. Pertandingan selesai.")
			break
		}

		// Menentukan pemenang
		if skor1 > skor2 {
			fmt.Printf("Hasil %d: %s menang\n", i, klub1)
			pemenang = append(pemenang, klub1)
		} else if skor2 > skor1 {
			fmt.Printf("Hasil %d: %s menang\n", i, klub2)
			pemenang = append(pemenang, klub2)
		} else {
			fmt.Printf("Hasil %d: Draw\n", i)
		}
	}

	// Menampilkan hasil akhir
	fmt.Println("\nDaftar klub yang memenangkan pertandingan:")
	if len(pemenang) > 0 {
		fmt.Println(strings.Join(pemenang, ", "))
	} else {
		fmt.Println("Tidak ada klub yang menang.")
	}
}
