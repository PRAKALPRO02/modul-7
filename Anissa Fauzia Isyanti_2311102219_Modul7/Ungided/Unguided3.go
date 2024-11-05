package main

import (
	"fmt"
)

func main() {
	var klubA, klubB string
	var skorA, skorB int
	var pemenang []string // Array untuk menyimpan hasil pemenang setiap pertandingan

	// Input nama klub yang bertanding
	fmt.Print("Klub A : ")
	fmt.Scanln(&klubA)
	fmt.Print("Klub B : ")
	fmt.Scanln(&klubB)

	//Input skor untuk setiap pertandingan
	pertandingan := 1
	for {
		fmt.Printf("Pertandingan %d : ", pertandingan)
		fmt.Scan(&skorA, &skorB)

		//Memeriksa validitas skor, jika negatif maka berhenti
		if skorA < 0 || skorB < 0 {
			break
		}

		//Menyimpan hasil pertandingan dalam array `pemenang`
		if skorA > skorB {
			pemenang = append(pemenang, klubA)
		} else if skorB > skorA {
			pemenang = append(pemenang, klubB)
		} else {
			pemenang = append(pemenang, "Draw")
		}

		pertandingan++
	}

	//Menampilkan daftar hasil pertandingan setelah semua input selesai
	fmt.Println("\nHasil Pertandingan:")
	for i, hasil := range pemenang {
		if hasil != "Draw" {
			fmt.Printf("Hasil %d : %s\n", i+1, hasil)
		} else {
			fmt.Printf("Hasil %d : Draw\n", i+1)
		}
	}

	//Menampilkan keterangan bahwa pertandingan sudah selesai
	fmt.Println("Pertandingan selesai")
}
