package main

import (
	"fmt"
	"math"
)

type point213 struct {
	x, y float64
}

func jarak(p, q point213) float64 {
	return math.Sqrt(math.Pow(p.x-q.x, 2) + math.Pow(p.y-q.y, 2))
}

func dalam(center point213, radius float64, p point213) bool {
	return jarak(center, p) <= radius
}

func main() {
	var cx1, cy1, r1 float64 
	var cx2, cy2, r2 float64 
	var px, py float64       

	fmt.Print("Masukkan cx1, cy1, dan r1: ")
	fmt.Scan(&cx1, &cy1, &r1)

	fmt.Print("Masukkan cx2, cy2, dan r2: ")
	fmt.Scan(&cx2, &cy2, &r2)

	fmt.Print("Masukkan px dan py: ")
	fmt.Scan(&px, &py)

	center1_213 := point213{cx1, cy1}
	center2_213 := point213{cx2, cy2}
	point213 := point213{px, py}

	Lngkrn1 := dalam(center1_213, r1, point213)
	Lngkrn2 := dalam(center2_213, r2, point213)

	if Lngkrn1 && Lngkrn2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if Lngkrn1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if Lngkrn2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
