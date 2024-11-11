package main

import "fmt"

func rekapPertandingan(klubA, klubB string, pertandingan int, pemenang *[]string) {
	var skorA, skorB int
	fmt.Printf("Pertandingan %d: ", pertandingan)
	fmt.Scanln(&skorA, &skorB)

	if skorA < 0 || skorB < 0 {
		return
	}

	if skorA > skorB {
		*pemenang = append(*pemenang, klubA)
	} else if skorA < skorB {
		*pemenang = append(*pemenang, klubB)
	} else {
		*pemenang = append(*pemenang, "Draw")
	}
	rekapPertandingan(klubA, klubB, pertandingan+1, pemenang)
}

func main() {
	var klubA, klubB string
	fmt.Print("Klub A: ")
	fmt.Scanln(&klubA)
	fmt.Print("Klub B: ")
	fmt.Scanln(&klubB)

	var pemenang []string
	rekapPertandingan(klubA, klubB, 1, &pemenang)

	for i, p := range pemenang {
		fmt.Printf("Hasil %d: %s\n", i+1, p)
	}
	fmt.Println("Pertandingan Selesai")
}
