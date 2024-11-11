package main

import "fmt"

func inputNamaKlub() (string, string) {
	var klubA, klubB string
	fmt.Print("Klub A: ")
	fmt.Scanln(&klubA)
	fmt.Print("Klub B: ")
	fmt.Scanln(&klubB)
	return klubA, klubB
}

func inputSkor(pertandingan int, klubA, klubB string) (int, int, bool) {
	var skorA, skorB int
	fmt.Printf("Pertandingan %d: ", pertandingan)
	fmt.Scanln(&skorA, &skorB)

	if skorA < 0 || skorB < 0 {
		return skorA, skorB, false
	}
	fmt.Printf("// %s %d sedangkan %s %d\n", klubA, skorA, klubB, skorB)
	return skorA, skorB, true
}

func tentukanPemenang(skorA, skorB int, klubA, klubB string) string {
	if skorA > skorB {
		return klubA
	} else if skorB > skorA {
		return klubB
	}
	return "Seri"
}

func tampilkanHasil(pemenang []string) {
	fmt.Println("Hasil pertandingan:")
	for i, hasil := range pemenang {
		fmt.Printf("Hasil %d: %s\n", i+1, hasil)
	}
	fmt.Println("Pertandingan selesai")
}

func main() {
	klubA, klubB := inputNamaKlub()

	var pemenang []string
	var pertandingan int = 1

	for {
		skorA, skorB, valid := inputSkor(pertandingan, klubA, klubB)

		if !valid {
			fmt.Println("Skor tidak valid")
			break
		}

		hasil := tentukanPemenang(skorA, skorB, klubA, klubB)
		pemenang = append(pemenang, hasil)
		pertandingan++

		if pertandingan > 12 {
			break
		}
	}

	tampilkanHasil(pemenang)
}
