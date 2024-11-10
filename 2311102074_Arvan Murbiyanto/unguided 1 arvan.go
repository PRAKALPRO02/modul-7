package main

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

func jarak(p, q Titik) float64 {
	return math.Sqrt(math.Pow(float64(p.x-q.x), 2) + math.Pow(float64(p.y-q.y), 2))
}

func didalam(c Lingkaran, p Titik) bool {
	return jarak(c.pusat, p) <= float64(c.radius)
}

func main() {
	var lingkaran1, lingkaran2 Lingkaran
	var titikSembarang Titik

	fmt.Print("Masukkan titik pusat (x, y) dan radius lingkaran 1: ")
	fmt.Scan(&lingkaran1.pusat.x, &lingkaran1.pusat.y, &lingkaran1.radius)

	fmt.Print("Masukkan titik pusat (x, y) dan radius lingkaran 2: ")
	fmt.Scan(&lingkaran2.pusat.x, &lingkaran2.pusat.y, &lingkaran2.radius)

	fmt.Print("Masukkan koordinat titik sembarang (x, y): ")
	fmt.Scan(&titikSembarang.x, &titikSembarang.y)

	diLingkaran1 := didalam(lingkaran1, titikSembarang)
	diLingkaran2 := didalam(lingkaran2, titikSembarang)

	if diLingkaran1 && diLingkaran2 {
		fmt.Println("Titik berada di dalam lingkaran 1 dan 2")
	} else if diLingkaran1 {
		fmt.Println("Titik berada di dalam lingkaran 1")
	} else if diLingkaran2 {
		fmt.Println("Titik berada di dalam lingkaran 2")
	} else {
		fmt.Println("Titik berada di luar lingkaran 1 dan 2")
	}
}
