package main

import (
	"fmt"
	"math"
)

type Titik struct {
	x, y float64
}
type Lingkaran struct {
	a, b float64
	Jari float64
}

func hitungjarak(titik Titik, titik_2 Lingkaran) float64 {
	hasil := math.Sqrt(math.Pow((titik.x-titik_2.a), 2) + math.Pow(titik.y-titik_2.b, 2))
	return hasil
}
func posisi(titik Titik, lingkaran Lingkaran) bool {
	hasil := hitungjarak(titik, lingkaran) <= lingkaran.Jari
	return hasil
}

func main() {
	var titik Titik
	var lingkaran_1, lingkaran_2 Lingkaran

	fmt.Print("Lingkaran 1 : ")
	fmt.Scan(&lingkaran_1.a, &lingkaran_1.b, &lingkaran_1.Jari)
	fmt.Print("Lingkaran 2 : ")
	fmt.Scan(&lingkaran_2.a, &lingkaran_2.b, &lingkaran_2.Jari)
	fmt.Print("Titik : ")
	fmt.Scan(&titik.x, &titik.y)

	if posisi(titik, lingkaran_1) == true && posisi(titik, lingkaran_2) == true {
		fmt.Print("Titik di dalam lingkaran 1 dan 2")
	} else if posisi(titik, lingkaran_1) == true {
		fmt.Print("Titik di dalam lingkaran 1")
	} else if posisi(titik, lingkaran_2) == true {
		fmt.Print("Titik di dalam lingkaran 2")
	} else {
		fmt.Print("Titik di luar lingkaran 1 dan 2")
	}

}
