package main

import (
	"fmt"
	"math"
)

type Titik225 struct {
	x, y int
}

type Lingkaran225 struct {
	center Titik225
	radius int
}

func jarak(p, q Titik225) float64 {
	return math.Sqrt(float64((p.x-q.x)*(p.x-q.x) + (p.y-q.y)*(p.y-q.y)))
}

func didalam(c Lingkaran225, p Titik225) bool {
	return jarak(c.center, p) <= float64(c.radius)
}

func main() {
	var cx1, cy1, r1 int
	var cx2, cy2, r2 int
	var tx, ty int

	fmt.Println("Masukkan titik pusat dan radius lingkaran 1 (cx1, cy1, r1):")
	fmt.Scan(&cx1, &cy1, &r1)
	fmt.Println("Masukkan titik pusat dan radius lingkaran 2 (cx2, cy2, r2):")
	fmt.Scan(&cx2, &cy2, &r2)
	fmt.Println("Masukkan koordinat titik sembarang (tx, ty):")
	fmt.Scan(&tx, &ty)

	lingkaran1 := Lingkaran225{Titik225{cx1, cy1}, r1}
	lingkaran2 := Lingkaran225{Titik225{cx2, cy2}, r2}
	titik := Titik225{tx, ty}

	diDalamLingkaran1 := didalam(lingkaran1, titik)
	diDalamLingkaran2 := didalam(lingkaran2, titik)

	if diDalamLingkaran1 && diDalamLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if diDalamLingkaran1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if diDalamLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
