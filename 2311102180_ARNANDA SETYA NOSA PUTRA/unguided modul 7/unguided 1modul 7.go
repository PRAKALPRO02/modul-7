package main

import (
	"fmt"
	"math"
)

type titik struct {
	x, y int
}

type lingkaran struct {
	pusat  titik
	radius int
}

func jarak(p, q titik) float64 {
	return math.Sqrt(float64((p.x-q.x)*(p.x-q.x) + (p.y-q.y)*(p.y-q.y)))
}

func didalam(c lingkaran, p titik) bool {
	return jarak(c.pusat, p) <= float64(c.radius)
}

func main() {
	var cx1, cy1, r1 int
	fmt.Println("Masukkan koordinat pusat dan radius lingkaran 1:")
	fmt.Scan(&cx1, &cy1, &r1)
	lingkaran1 := lingkaran{pusat: titik{x: cx1, y: cy1}, radius: r1}

	var cx2, cy2, r2 int
	fmt.Println("Masukkan koordinat pusat dan radius lingkaran 2:")
	fmt.Scan(&cx2, &cy2, &r2)
	lingkaran2 := lingkaran{pusat: titik{x: cx2, y: cy2}, radius: r2}

	var x, y int
	fmt.Println("Masukkan koordinat titik sembarang:")
	fmt.Scan(&x, &y)
	titikSembarang := titik{x: x, y: y}

	diLingkaran1 := didalam(lingkaran1, titikSembarang)
	diLingkaran2 := didalam(lingkaran2, titikSembarang)

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
