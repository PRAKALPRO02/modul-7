package main

import (
	"fmt"
	"math"
)

type titik struct {
	a, b float64
}
type lingkaran struct {
	x, y, r float64
}

func main() {
	var l1, l2 lingkaran
	var x titik
	fmt.Scan(&l1.x, &l1.y, &l1.r, &l2.x, &l2.y, &l2.r, &x.a, &x.b)
	dalam1 := didalam(l1, x)
	dalam2 := didalam(l2, x)

	if dalam1 && dalam2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dalam1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dalam2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}

func jarak(a lingkaran, c titik) float64 {
	return math.Sqrt(math.Pow(c.a-a.x, 2) + math.Pow(c.b-a.y, 2))
}

func didalam(l lingkaran, x titik) bool {
	return jarak(l, x) <= l.r
}
