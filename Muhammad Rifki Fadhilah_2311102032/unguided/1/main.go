package main

import (
	"fmt"
	"math"
)

type titik int
type lingkaran int

func jarak(a, b, c, d titik) float64 {
	return math.Sqrt(math.Pow(float64(a-c), 2) + math.Pow(float64(b-d), 2))
}

func didalam(cx, cy lingkaran, r lingkaran, x, y titik) bool {
	jarak_titik := jarak(x, y, titik(cx), titik(cy))
	return jarak_titik < float64(r)
}

func main() {
	var cx1, cy1, r1, cx2, cy2, r2, x, y titik

	fmt.Scanln(&cx1, &cy1, &r1)
	fmt.Scanln(&cx2, &cy2, &r2)
	fmt.Scanln(&x, &y)

	result1 := didalam(lingkaran(cx1), lingkaran(cy1), lingkaran(r1), x, y)
	result2 := didalam(lingkaran(cx2), lingkaran(cy2), lingkaran(r2), x, y)

	if result1 && result2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if result1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if result2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
