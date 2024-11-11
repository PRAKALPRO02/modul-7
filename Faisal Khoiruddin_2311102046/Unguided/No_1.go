package main

import (
	"fmt"
	"math"
)

type Titik struct {
	x, y int
}

type Lingkaran struct {
	titik  Titik
	radius int
}

func jarak(t1, t2 Titik) float64 {
	return math.Sqrt(float64((t2.x-t1.x)*(t2.x-t1.x) + (t2.y-t1.y)*(t2.y-t1.y)))
}

func posisiTitikTerhadapLingkaran(l Lingkaran, t Titik) bool {
	return jarak(l.titik, t) <= float64(l.radius)
}

func main() {
	var x1, y1, r1, x2, y2, r2, xt, yt int
	fmt.Scan(&x1, &y1, &r1)
	lingkaran1 := Lingkaran{Titik{x1, y1}, r1}
	fmt.Scan(&x2, &y2, &r2)
	lingkaran2 := Lingkaran{Titik{x2, y2}, r2}
	fmt.Scan(&xt, &yt)
	titik := Titik{xt, yt}

	dalamLingkaran1 := posisiTitikTerhadapLingkaran(lingkaran1, titik)
	dalamLingkaran2 := posisiTitikTerhadapLingkaran(lingkaran2, titik)

	if dalamLingkaran1 && dalamLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dalamLingkaran1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dalamLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
