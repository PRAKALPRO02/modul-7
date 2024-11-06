package main

import (
	"fmt"
	"math"
)

type Titik struct {
	x, y float64
}

type Lingkaran struct {
	pusat  Titik
	radius float64
}

func jarak(p, q Titik) float64 {
	return math.Sqrt(math.Pow(p.x-q.x, 2) + math.Pow(p.y-q.y, 2))
}

func didalam(c Lingkaran, p Titik) bool {
	return jarak(c.pusat, p) <= c.radius
}

func main() {
	var l1, l2 Lingkaran
	var p Titik

	fmt.Print("Masukkan koordinat pusat dan radius lingkaran 1: ")
	fmt.Scan(&l1.pusat.x, &l1.pusat.y, &l1.radius)
	fmt.Print("Masukkan koordinat pusat dan radius lingkaran 2: ")
	fmt.Scan(&l2.pusat.x, &l2.pusat.y, &l2.radius)
	fmt.Print("Masukkan koordinat titik sembarang: ")
	fmt.Scan(&p.x, &p.y)

	inL1 := didalam(l1, p)
	inL2 := didalam(l2, p)

	if inL1 && inL2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if inL1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if inL2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
