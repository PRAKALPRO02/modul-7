package main

import "fmt"

type nama struct {
	Nama string
	Skor int
}

func main() {
	var klubA, klubB nama
	var pemenang []string
	var tanding int = 1

	fmt.Print("Masukkan nama Klub pertama: ")
	fmt.Scanln(&klubA.Nama)
	fmt.Print("Masukkan nama Klub kedua: ")
	fmt.Scanln(&klubB.Nama)

	for {
		fmt.Printf("Pertandingan %v: ", tanding)
		fmt.Scanln(&klubA.Skor, &klubB.Skor)
		tanding++

		if klubA.Skor < 0 || klubB.Skor < 0 {
			break
		}

		if klubA.Skor > klubB.Skor {
			pemenang = append(pemenang, klubA.Nama)
		} else if klubB.Skor == klubA.Skor {
			pemenang = append(pemenang, "Draw")
		} else if klubB.Skor > klubA.Skor {
			pemenang = append(pemenang, klubB.Nama)
		}
	}

	fmt.Println("Daftar klub yang memenangkan pertandingan:")
	for _, nama := range pemenang {
		fmt.Println(nama)
	}
}
