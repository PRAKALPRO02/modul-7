package main

import (
	"fmt"
)

func main() {
	var klubA, klubB string
	fmt.Print("Masukkan nama Klub A: ")
	fmt.Scanln(&klubA)
	fmt.Print("Masukkan nama Klub B: ")
	fmt.Scanln(&klubB)
	var pemenang []string
	var pertandingan int
	for i := 1; ; i++ {
		var skorA, skorB int
		fmt.Printf("Masukkan skor pertandingan %d (%s vs %s): ", i, klubA,
			klubB)
		fmt.Scan(&skorA, &skorB)
		if skorA < 0 || skorB < 0 {
			fmt.Println("Skor tidak valid, menghentikan input...")
			pertandingan = i - 1
			break
		}
		if skorA > skorB {
			pemenang = append(pemenang, fmt.Sprintf("Hasil %d: %s menang", i,
				klubA))
		} else if skorB > skorA {
			pemenang = append(pemenang, fmt.Sprintf("Hasil %d: %s menang", i,
				klubB))
		} else {
			pemenang = append(pemenang, fmt.Sprintf("Hasil %d: Draw", i))
		}
	}
	fmt.Println("Hasil pertandingan:")
	for i := 0; i < pertandingan; i++ {
		fmt.Println(pemenang[i])
	}
	fmt.Println("Pertandingan selesai")
}
