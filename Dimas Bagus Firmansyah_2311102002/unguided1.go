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
	radius int
}

// Fungsi untuk menghitung jarak antara dua titik
func jarak(a, b Titik) float64 {
	return math.Sqrt(math.Pow(float64(a.x-b.x), 2) + math.Pow(float64(a.y-b.y), 2))
}

// Fungsi untuk menentukan apakah titik berada di dalam lingkaran
func dalamLingkaran(t Titik, l Lingkaran) bool {
	return jarak(t, l.pusat) <= float64(l.radius)
}

func main() {
	// Input data lingkaran 1
	var cx1, cy1, r1 int
	fmt.Scan(&cx1, &cy1, &r1)
	lingkaran1 := Lingkaran{Titik{cx1, cy1}, r1}

	// Input data lingkaran 2
	var cx2, cy2, r2 int
	fmt.Scan(&cx2, &cy2, &r2)
	lingkaran2 := Lingkaran{Titik{cx2, cy2}, r2}

	// Input data titik sembarang
	var x, y int
	fmt.Scan(&x, &y)
	titik := Titik{x, y}

	// Cek posisi titik terhadap kedua lingkaran
	dalamL1 := dalamLingkaran(titik, lingkaran1)
	dalamL2 := dalamLingkaran(titik, lingkaran2)

	// Output berdasarkan posisi titik
	if dalamL1 && dalamL2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dalamL1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dalamL2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
