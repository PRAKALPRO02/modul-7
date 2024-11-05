package main

import (
	"fmt"
	"math"
)

type Titik struct {
	x, y int
}

type Lingkaran struct {
	pusat Titik
	r     int
}

func cekTitikDiDalamLingkaran(t Titik, l Lingkaran) bool {
	jarak := math.Sqrt(math.Pow(float64(t.x-l.pusat.x), 2) + math.Pow(float64(t.y-l.pusat.y), 2))
	return jarak <= float64(l.r)
}

func main() {
	
	var x1, y1, r1 int
	fmt.Print("Masukkan lingkaran 1 (x y r): ")
	fmt.Scan(&x1, &y1, &r1)
	lingkaran1 := Lingkaran{Titik{x1, y1}, r1}

	var x2, y2, r2 int
	fmt.Print("Masukkan lingkaran 2 (x y r): ")
	fmt.Scan(&x2, &y2, &r2)
	lingkaran2 := Lingkaran{Titik{x2, y2}, r2}

	var x, y int
	fmt.Print("Masukkan titik (x y): ")
	fmt.Scan(&x, &y)
	titik := Titik{x, y}

	diLingkaran1 := cekTitikDiDalamLingkaran(titik, lingkaran1)
	diLingkaran2 := cekTitikDiDalamLingkaran(titik, lingkaran2)

	if diLingkaran1 && diLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if diLingkaran1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if diLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
