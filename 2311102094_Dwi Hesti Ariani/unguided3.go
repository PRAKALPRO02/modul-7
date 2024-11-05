package main

import (
	"fmt"
)

func main() {
	var klubA, klubB string
	var skorA, skorB int
	var winners []string
	var results []string

	fmt.Print("Klub A : ")
	fmt.Scanln(&klubA)

	fmt.Print("Klub B : ")
	fmt.Scanln(&klubB)

	pertandinganKe := 1

	for {
		fmt.Printf("Pertandingan %d (skor A B): ", pertandinganKe)
		fmt.Scanln(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			winners = append(winners, klubA)
			results = append(results, fmt.Sprintf("Pertandingan %d: %s %d - %d %s (Pemenang: %s)", pertandinganKe, klubA, skorA, skorB, klubB, klubA))
		} else if skorB > skorA {
			winners = append(winners, klubB)
			results = append(results, fmt.Sprintf("Pertandingan %d: %s %d - %d %s (Pemenang: %s)", pertandinganKe, klubA, skorA, skorB, klubB, klubB))
		} else {
			results = append(results, fmt.Sprintf("Pertandingan %d: %s %d - %d %s (Hasil Seri)", pertandinganKe, klubA, skorA, skorB, klubB))
		}

		pertandinganKe++
	}

	fmt.Println("Pertandingan selesai")
	fmt.Println("Hasil setiap pertandingan:")
	for _, result := range results {
		fmt.Println(result)
	}
}
