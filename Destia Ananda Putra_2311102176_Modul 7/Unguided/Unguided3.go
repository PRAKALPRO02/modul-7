package main

// Destia Ananda Putra
// 2311102176

import (
	"fmt"
)

func main() {

	var klubA, klubB string
	fmt.Print("Klub A : ")
	fmt.Scanln(&klubA)
	fmt.Print("Klub B : ")
	fmt.Scanln(&klubB)

	var hasil []string

	var skorA, skorB int
	pertandingan := 1

	for {
		fmt.Printf("Pertandingan %d : ", pertandingan)
		fmt.Scanln(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			fmt.Println("Pertandingan selesai")
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

	for i, result := range hasil {
		fmt.Printf("Hasil %d : %s\n", i+1, result)
	}

	fmt.Println("Pertandingan selesai")
}
