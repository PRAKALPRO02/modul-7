package main

import (
	"fmt"
)

func main() {
	var klubA, klubB string

	fmt.Print("Klub A : ")
	fmt.Scanln(&klubA)
	fmt.Print("Klub B : ")
	fmt.Scanln(&klubB)

	var pertandingan []string

	fmt.Println()
	for i := 1; ; i++ {
		var skorA, skorB int
		fmt.Printf("Pertandingan %d : ", i)
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			pertandingan = append(pertandingan, fmt.Sprintf("Hasil %d : %s", i, klubA))
		} else if skorB > skorA {
			pertandingan = append(pertandingan, fmt.Sprintf("Hasil %d : %s", i, klubB))
		} else {
			pertandingan = append(pertandingan, fmt.Sprintf("Hasil %d : Draw", i))
		}
	}

	fmt.Println()
	for _, hasil := range pertandingan {
		fmt.Println(hasil)
	}

	fmt.Println("Pertandingan selesai")
}
