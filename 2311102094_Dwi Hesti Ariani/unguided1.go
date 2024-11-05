package main

import (
	"fmt"
	"math"
)

type titik struct {
	x int
	y int
}

type lingkaran struct {
	titikPusat titik
	radius     int
}

func jarak(p, q titik) float64 {
	return math.Sqrt(float64((p.x-q.x)*(p.x-q.x) + (p.y-q.y)*(p.y-q.y)))
}

func didalam(c lingkaran, p titik) bool {
	return jarak(c.titikPusat, p) <= float64(c.radius)
}

func main() {
	var cx1, cy1, r1, cx2, cy2, r2, zx, zy int
	fmt.Scanln(&cx1, &cy1, &r1)
	fmt.Scanln(&cx2, &cy2, &r2)
	fmt.Scanln(&zx, &zy)

	l1 := lingkaran{titik{cx1, cy1}, r1}
	l2 := lingkaran{titik{cx2, cy2}, r2}
	p := titik{zx, zy}

	if didalam(l1, p) && didalam(l2, p) {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if didalam(l1, p) {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if didalam(l2, p) {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
