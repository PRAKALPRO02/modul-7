package main

import (
	"fmt"
)

func main() {
	var klubA, klubB string
	var skorA, skorB int
	var hasil []string

	fmt.Print("Masukkan nama Klub A: ")
	fmt.Scan(&klubA)
	//fahrur059
	fmt.Print("Masukkan nama Klub B: ")
	fmt.Scan(&klubB)
	fmt.Print("\n")

	for {
		fmt.Printf("Masukkan skor %s dan %s : ", klubA, klubB)
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			fmt.Println("Program dihentikan.")
			break
		}

		if skorA > skorB {
			hasil = append(hasil, klubA)
			fmt.Printf("Hasil: %s menang\n", klubA)
		} else if skorB > skorA {
			//fahrur059
			hasil = append(hasil, klubB)
			fmt.Printf("Hasil: %s menang\n", klubB)
		} else {
			fmt.Println("Hasil: Draw")
		}
	}

	fmt.Println("\nDaftar klub yang memenangkan pertandingan:")
	for _, klub := range hasil {
		fmt.Println(klub)
	}
}
