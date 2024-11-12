package main

import "fmt"

type pertandingan struct {
	klub1 string
	klub2 string
	hasil []string
}

func main() {
	var pertandingan pertandingan
	var skor_k1, skor_k2 int
	fmt.Print("Klub A: ")
	fmt.Scanln(&pertandingan.klub1)
	fmt.Print("Klub B: ")
	fmt.Scanln(&pertandingan.klub2)
	for i := 1; i > 0; i++ {
		fmt.Printf("Pertandingan %v: ", i)
		fmt.Scanln(&skor_k1, &skor_k2)
		if skor_k1 < 0 || skor_k2 < 0 {
			break
		} else if skor_k1 > skor_k2 {
			pertandingan.hasil = append(pertandingan.hasil, pertandingan.klub1)
		} else if skor_k1 < skor_k2 {
			pertandingan.hasil = append(pertandingan.hasil, pertandingan.klub2)
		} else {
			pertandingan.hasil = append(pertandingan.hasil, "Draw")
		}
	}
	for i := 0; i < len(pertandingan.hasil); i++ {
		fmt.Printf("Hasil %v: %v\n", i+1, pertandingan.hasil[i])
	}
	fmt.Println("Pertandingan selesai")
}