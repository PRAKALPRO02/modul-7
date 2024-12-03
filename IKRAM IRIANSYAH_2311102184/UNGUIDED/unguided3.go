package main

import (
	"fmt"
)

type Klub struct {
	klubA, klubB string
}

func input_klub() Klub {
	var klub Klub
	fmt.Print("Klub A : ")
	fmt.Scan(&klub.klubA)
	fmt.Print("Klub B : ")
	fmt.Scan(&klub.klubB)
	return klub
}

func input_skor(pertandingan int) (int, int) {
	var skorA, skorB int
	fmt.Printf("Pertandingan %d (Skor A, Skor B) : ", pertandingan)
	fmt.Scan(&skorA, &skorB)
	return skorA, skorB
}

func hasil_pemenang(skorA, skorB int, klub Klub) string {
	if skorA > skorB {
		return klub.klubA
	} else if skorB > skorA {
		return klub.klubB
	}
	return "Draw"
}

func hasil_pertandingan(pemenang []string) {
	fmt.Println("\nHasil pertandingan :")
	for i, winner := range pemenang {
		fmt.Printf("Hasil %d : %s\n", i+1, winner)
	}
	fmt.Println("Pertandingan selesai.")
}

func main() {
	// Input nama klub
	klub := input_klub()

	var pemenang []string
	pertandingan := 1

	for {
		skorA, skorB := input_skor(pertandingan)

		if skorA < 0 || skorB < 0 {
			break
		}

		winner := hasil_pemenang(skorA, skorB, klub)
		pemenang = append(pemenang, winner)

		pertandingan++
	}
	hasil_pertandingan(pemenang)
}
