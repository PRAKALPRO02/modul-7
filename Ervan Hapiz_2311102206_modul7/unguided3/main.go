package main

import "fmt"

type nama struct {
	Nama string
	Skor int
}

func main() {
	var nama1, nama2 nama
	var pemenang []string
	var tanding int = 1

	fmt.Print("Masukkan nama nama pertama: ")
	fmt.Scanln(&nama1.Nama)
	fmt.Print("Masukkan nama nama kedua: ")
	fmt.Scanln(&nama2.Nama)

	for {
		fmt.Printf("pertandingan %v: ", tanding)
		fmt.Scanln(&nama1.Skor, &nama2.Skor)
		tanding++

		if nama1.Skor < 0 || nama2.Skor < 0 {
			break
		}

		if nama1.Skor > nama2.Skor {
			pemenang = append(pemenang, nama1.Nama)
		} else if nama2.Skor == nama1.Skor {
			pemenang = append(pemenang, "DRAW")
		} else if nama2.Skor > nama1.Skor {
			pemenang = append(pemenang, nama2.Nama)
		}
	}

	fmt.Println("Daftar klub yang memenangkan pertandingan:")
	for _, nama := range pemenang {
		fmt.Println(nama)
	}
}
