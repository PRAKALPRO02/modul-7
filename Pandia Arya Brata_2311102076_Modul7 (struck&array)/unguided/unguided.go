package main

import (
	"fmt"
	"math"
)

type Titik struct {
	X int
	Y int
}

type Circle struct {
	Center Titik
	Radius int
}

func jarak(a, b Titik) float64 {
	return math.Sqrt(float64((a.X-b.X)*(a.X-b.X) + (a.Y-b.Y)*(a.Y-b.Y)))
}

func priksa_titik(titik Titik, circle1, circle2 Circle) string {
	dist1 := jarak(titik, circle1.Center)
	dist2 := jarak(titik, circle2.Center)

	if dist1 <= float64(circle1.Radius) && dist2 <= float64(circle2.Radius) {
		return "Titik di dalam lingkaran 1 dan 2"
	} else if dist1 <= float64(circle1.Radius) {
		return "Titik di dalam lingkaran 1"
	} else if dist2 <= float64(circle2.Radius) {
		return "Titik di dalam lingkaran 2"
	} else {
		return "Titik di luar lingkaran 1 dan 2"
	}
}

func main() {
	var (
		circle1 Circle
		circle2 Circle
		titik   Titik
	)

	fmt.Println("masukan kordinat titik pusat satu dan radius : ")
	fmt.Scanln(&circle1.Center.X, &circle1.Center.Y, &circle1.Radius)

	fmt.Println("masukan kordinat titik pusat dua dan radius : ")
	fmt.Scanln(&circle2.Center.X, &circle2.Center.Y, &circle2.Radius)

	fmt.Println("Masukkan koordinat titik sembarang (x, y):")
	fmt.Scanln(&titik.X, &titik.Y)

	result := priksa_titik(titik, circle1, circle2)
	fmt.Println(result)
}
