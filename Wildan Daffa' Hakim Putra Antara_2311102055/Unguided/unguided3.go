package main

import "fmt"

type Pertandingan struct {
	klubA             string
	klubB             string
	hasilPertandingan []string
}

func hitung(dataClub Pertandingan, skor1, skor2 int) string {
	if skor1 > skor2 {
		return dataClub.klubA
	} else if skor1 == skor2 {
		return "Draw"
	} else {
		return dataClub.klubB
	}
}

func showHasil(dataClub Pertandingan) {
	for i, data := range dataClub.hasilPertandingan {
		fmt.Printf("Hasil %d : %s\n", i+1, data)
	}
	fmt.Println("Pertandingan selesai")
}

func main() {
	var skor1, skor2, count int
	var dataTanding Pertandingan
	fmt.Print("Klub A : ")
	fmt.Scan(&dataTanding.klubA)
	fmt.Print("Klub B : ")
	fmt.Scan(&dataTanding.klubB)
	for {
		fmt.Printf("Pertandingan %d : ", count+1)
		fmt.Scan(&skor1, &skor2)
		if skor1 >= 0 && skor2 >= 0 {
			dataTanding.hasilPertandingan = append(dataTanding.hasilPertandingan, hitung(dataTanding, skor1, skor2))
			count++
		} else {
			break
		}

	}
	showHasil(dataTanding)
}
