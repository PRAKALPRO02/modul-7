package main

import "fmt"

func main() {
	var klubA, klubB string
	var skorA, skorB int
	var pemenang []string

	fmt.Print("Klub A: ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B: ")
	fmt.Scan(&klubB)

	pertandingan := 1
	for {
		fmt.Printf("Pertandingan %d: ", pertandingan)
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			pemenang = append(pemenang, klubA)
		} else if skorB > skorA {
			pemenang = append(pemenang, klubB) 
		} else {
			pemenang = append(pemenang, "Draw")
		}

		pertandingan++
	}

	for i, winner := range pemenang {
		if winner != "Draw" {
			fmt.Printf("Pertandingan %d: %s\n", i+1, winner)
		} else {
			fmt.Printf("Pertandingan %d: Draw\n", i+1)
		}
	}
  fmt.Print("Pertandingan selesai")
}
