package main

import (
	"fmt"
)

func main() {
	var A213, B213 string
	var skorA, skorB int
	var pemenang213 []string
	var skorAList, skorBList []int

	fmt.Print("Klub A: ")
	fmt.Scanln(&A213)
	fmt.Print("Klub B: ")
	fmt.Scanln(&B213)

	pertandingan := 1
	for pertandingan <= 9 {
		fmt.Printf("Pertandingan %d:\n", pertandingan)
		fmt.Printf("Skor %s: ", A213)
		fmt.Scan(&skorA)
		fmt.Printf("Skor %s: ", B213)
		fmt.Scan(&skorB)

		if skorA < 0 || skorB < 0 {
			break
		}
		skorAList = append(skorAList, skorA)
		skorBList = append(skorBList, skorB)
		if skorA > skorB {
			pemenang213 = append(pemenang213, A213)
		} else if skorB > skorA {
			pemenang213 = append(pemenang213, B213)
		} else {
			pemenang213 = append(pemenang213, "Draw")
		}

		pertandingan++
	}

	fmt.Println("\nHasil Skor dan Pertandingan:")
	for i := 0; i < len(skorAList); i++ {
		fmt.Printf("Pertandingan %d : %d  %d\n", i+1, skorAList[i], skorBList[i])
	}

	fmt.Println("\nHasil Pertandingan:")
	for i, hasil := range pemenang213 {
		fmt.Printf("Hasil %d : %s\n", i+1, hasil)
	}
}
