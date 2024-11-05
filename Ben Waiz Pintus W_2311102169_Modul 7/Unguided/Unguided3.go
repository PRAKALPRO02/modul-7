// Ben Waiz Pintus W
// 2311102169
package main

import "fmt"

func main() {
	var klubA, klubB string
	var skorA, skorB int
	var pemenang []string

	fmt.Print("Klub A: ")
	fmt.Scanln(&klubA)
	fmt.Print("Klub B: ")
	fmt.Scanln(&klubB)

	pertandingan := 1
	for {
		fmt.Printf("Pertandingan %d: ", pertandingan)
		fmt.Scanln(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			pemenang = append(pemenang, klubA)
		} else if skorA < skorB {
			pemenang = append(pemenang, klubB)
		} else {
			pemenang = append(pemenang, "Draw")
		}

		pertandingan++
	}

	fmt.Println("Pertandingan selesai")
	for i, p := range pemenang {
		fmt.Printf("Hasil %d: %s\n", i+1, p)
	}
}
