package main

// NAMA : CHRIST DANIEL SANTOSO
// NIM : 2311102305

import (
	"fmt"
	"math"
)

type Titik struct {
	X, Y int
}

type Lingkaran struct {
	Pusat    Titik
	JariJari int
}

func hitungJarak(t Titik, l Lingkaran) float64 {
	return math.Sqrt(math.Pow(float64(t.X-l.Pusat.X), 2) + math.Pow(float64(t.Y-l.Pusat.Y), 2))
}

func titikDiDalamLingkaran(t Titik, l Lingkaran) bool {
	return hitungJarak(t, l) <= float64(l.JariJari)
}

func main() {
\	var x1, y1, r1 int
	fmt.Print("Masukkan untuk lingkaran 1: ")
	fmt.Scan(&x1, &y1, &r1)
	lingkaran1 := Lingkaran{Pusat: Titik{X: x1, Y: y1}, JariJari: r1}

	var x2, y2, r2 int
	fmt.Print("Masukkan untuk lingkaran 2: ")
	fmt.Scan(&x2, &y2, &r2)
	lingkaran2 := Lingkaran{Pusat: Titik{X: x2, Y: y2}, JariJari: r2}

	var x, y int
	fmt.Print("Masukkan untuk titik sembarang: ")
	fmt.Scan(&x, &y)
	titik := Titik{X: x, Y: y}

	diLingkaran1 := titikDiDalamLingkaran(titik, lingkaran1)
	diLingkaran2 := titikDiDalamLingkaran(titik, lingkaran2)

	switch {
	case diLingkaran1 && diLingkaran2:
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	case diLingkaran1:
		fmt.Println("Titik di dalam lingkaran 1")
	case diLingkaran2:
		fmt.Println("Titik di dalam lingkaran 2")
	default:
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
