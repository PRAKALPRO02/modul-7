package main

import "fmt"

func main(){
	var klubA, klubB string
	var skorA, skorB int
	var menang[] string
	var hasil[] string

	fmt.Print("Klub A : ")
	fmt.Scanln(&klubA)
	fmt.Print("Klub B : ")
	fmt.Scanln(&klubB)


	pertandingan := 1
	for {
		fmt.Printf("Pertandingan %d : ", pertandingan)
		fmt.Scanln(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			//fmt.Print("Skor tidak valid")
			break
		}

		if skorA > skorB {
			menang = append(menang, klubA)
			hasil = append(hasil, fmt.Sprintf("Hasil %d: %s ", pertandingan,  klubA))

		} else if skorB > skorA {
			menang = append(menang, klubB)
			hasil = append(hasil, fmt.Sprintf("Hasil %d: %s ", pertandingan, klubB))

		} else {
			hasil = append(hasil, fmt.Sprintf("Hasil %d: Draw ", pertandingan))
		}
		pertandingan++
	}
	fmt.Println()
	for _, klub := range hasil {
		fmt.Println(klub)
	}
	fmt.Println("\nPertandingan Selesai!")
}