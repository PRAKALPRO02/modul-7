package main

import "fmt"

func main() {
	var klubA, klubB string
	var skorA, skorB int
	var hasil []string

	fmt.Println("Masukkan nama klub A:")
	fmt.Scanln(&klubA)
	fmt.Println("Masukkan nama klub B:")
	fmt.Scanln(&klubB)

	for {
		fmt.Println("Masukkan skor klub A:")
		fmt.Scanln(&skorA)
		fmt.Println("Masukkan skor klub B:")
		fmt.Scanln(&skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			hasil = append(hasil, klubA)
		} else if skorB > skorA {
			hasil = append(hasil, klubB)
		} else {
			hasil = append(hasil, "Draw")
		}
	}

	fmt.Println("\nHasil pertandingan:")
	for i, v := range hasil {
		fmt.Printf("Pertandingan %d: %s\n", i+1, v)
	}
}

// Zahrina Antika Malahati_2311102109
