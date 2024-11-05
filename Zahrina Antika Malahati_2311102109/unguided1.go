package main

import (
	"fmt"
	"math"
)

// Titik mepresentasikan sebuah titik dalam ruang 2D
type Titik struct {
	X int
	Y int
}

// Lingkaran mepresentasikan sebuah lingkaran dengan titik pusat dan jari-jari
type Lingkaran struct {
	Pusat  Titik
	Radius int
}

// Jarak menghitung anatara dua titik
func Jarak(p, q Titik) float64 {
	return math.Sqrt(math.Pow(float64(p.X-q.X), 2) + math.Pow(float64(p.Y-q.Y), 2))
}

// Didalam memeriksa apakah sebuah titik berada dalam lingkaran
func Didalam(c Lingkaran, p Titik) bool {
	return Jarak(c.Pusat, p) <= float64(c.Radius)
}

func main() {
	var l1, l2 Lingkaran
	var p Titik

	// Input untuk lingakaran 1
	fmt.Println("Masukkan koordinat titik pusat dan radius lingkaran 1 (x y r):")
	fmt.Scanln(&l1.Pusat.X, &l1.Pusat.Y, &l1.Radius)

	// Input untuk lingkaran 2
	fmt.Println("Masukkan koordinat titik pusat dan radius lingkaran 2 (x y r):")
	fmt.Scanln(&l2.Pusat.X, &l2.Pusat.Y, &l2.Radius)

	// Input untuk titik yang akan diperiksa
	fmt.Println("Masukkan koordinat titik (x y):")
	fmt.Scanln(&p.X, &p.Y)

	// Menentukan posisi titik terhadap lingkaran
	if Didalam(l1, p) && Didalam(l2, p) {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if Didalam(l1, p) {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if Didalam(l2, p) {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}

// Zahrina Antika Malahati_2311102109
