package main

import (
	"fmt"
	"math"
)

type Titik struct {
	x, y float64
}

type Lingkaran struct {
	titik  Titik
	radius float64
}

func jarak(a, b Titik) float64 {
	return math.Sqrt(math.Pow(a.x-b.x, 2) + math.Pow(a.y-b.y, 2))
}

func diDalamLingkaran(t Titik, l Lingkaran) bool {
	return jarak(t, l.titik) <= l.radius
}

func main() {
	var l1, l2 Lingkaran
	var t Titik

	fmt.Print("Masukkan: ")
	fmt.Scan(&l1.titik.x, &l1.titik.y, &l1.radius)

	fmt.Print("Masukkan: ")
	fmt.Scan(&l2.titik.x, &l2.titik.y, &l2.radius)

	fmt.Print("Masukkan: ")
	fmt.Scan(&t.x, &t.y)

	dalamL1 := diDalamLingkaran(t, l1)
	dalamL2 := diDalamLingkaran(t, l2)

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
