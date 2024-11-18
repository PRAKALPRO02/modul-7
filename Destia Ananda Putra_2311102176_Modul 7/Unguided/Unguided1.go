package main

// Destia Ananda Putra
// 2311102176

import (
	"fmt"
	"math"
)

type Titik struct {
	x, y int
}

type Lingkaran struct {
	pusat  Titik
	radius int
}

func hitungJarak(t Titik, l Lingkaran) float64 {
	return math.Sqrt(math.Pow(float64(t.x-l.pusat.x), 2) + math.Pow(float64(t.y-l.pusat.y), 2))
}

func apakahDiDalamLingkaran(jarak float64, radius int) bool {
	return jarak <= float64(radius)
}

func main() {
	var x1, y1, r1 int
	fmt.Print("Masukkan koordinat x1, y1 dan radius r1 lingkaran pertama: ")
	fmt.Scan(&x1, &y1, &r1)
	lingkaran1 := Lingkaran{Titik{x1, y1}, r1}

	var x2, y2, r2 int
	fmt.Print("Masukkan koordinat x2, y2 dan radius r2 lingkaran kedua: ")
	fmt.Scan(&x2, &y2, &r2)
	lingkaran2 := Lingkaran{Titik{x2, y2}, r2}

	var x, y int
	fmt.Print("Masukkan koordinat x dan y titik sembarang: ")
	fmt.Scan(&x, &y)
	titik := Titik{x, y}

	jarakKeLingkaran1 := hitungJarak(titik, lingkaran1)
	jarakKeLingkaran2 := hitungJarak(titik, lingkaran2)

	dalamLingkaran1 := apakahDiDalamLingkaran(jarakKeLingkaran1, lingkaran1.radius)
	dalamLingkaran2 := apakahDiDalamLingkaran(jarakKeLingkaran2, lingkaran2.radius)

	if dalamLingkaran1 && dalamLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dalamLingkaran1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dalamLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
