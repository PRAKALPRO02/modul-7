package main

import (
	"fmt"
	"math"
)

type Titik struct {
	x, y int
}

type Lingkaran struct {
	radius     float64
	titikPusat Titik
}

func getJarak(lingkaran Lingkaran, titikSembarang Titik) float64 {
	return math.Sqrt(math.Pow(float64(lingkaran.titikPusat.x)-float64(titikSembarang.x), 2) + math.Pow(float64(lingkaran.titikPusat.y)-float64(titikSembarang.y), 2))
}

func cekDidalam(lingkaran Lingkaran, jarak float64) bool {
	return lingkaran.radius >= jarak
}

func main() {
	var lingkaran1, lingkaran2 Lingkaran
	var titikSembarang Titik
	var didalam1, didalam2 bool
	fmt.Scanln(&lingkaran1.titikPusat.x, &lingkaran1.titikPusat.y, &lingkaran1.radius)
	fmt.Scanln(&lingkaran2.titikPusat.x, &lingkaran2.titikPusat.y, &lingkaran2.radius)
	fmt.Scanln(&titikSembarang.x, &titikSembarang.y)
	didalam1 = cekDidalam(lingkaran1, getJarak(lingkaran1, titikSembarang))
	didalam2 = cekDidalam(lingkaran2, getJarak(lingkaran2, titikSembarang))
	switch {
	case didalam2 && didalam1:
		fmt.Print("Titik di dalam lingkaran 1 dan 2")
	case didalam1:
		fmt.Print("Titik di dalam lingkaran 1")
	case didalam2:
		fmt.Print("Titik di dalam lingkaran 2")
	default:
		fmt.Print("Titik di luar lingkaran 1 dan 2")
	}
}
