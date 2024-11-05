package main

import (
	"fmt"
	"math"
)

type Titik234 struct {
	x, y int
}

type Lingkaran234 struct {
	pusat  Titik234
	radius int
}

func jarak234(p, q Titik234) float64 {
	return math.Sqrt(float64((p.x-q.x)*(p.x-q.x) + (p.y-q.y)*(p.y-q.y)))
}

func diDalam234(c Lingkaran234, p Titik234) bool {
	return jarak234(c.pusat, p) <= float64(c.radius)
}

func main() {
	var cx1234, cy1234, r1234, cx2234, cy2234, r2234, x234, y234 int

	fmt.Println("Masukkan koordinat pusat dan radius lingkaran 1:")
	fmt.Scan(&cx1234, &cy1234, &r1234)
	lingkaran1234 := Lingkaran234{Titik234{cx1234, cy1234}, r1234}

	fmt.Println("Masukkan koordinat pusat dan radius lingkaran 2:")
	fmt.Scan(&cx2234, &cy2234, &r2234)
	lingkaran2234 := Lingkaran234{Titik234{cx2234, cy2234}, r2234}

	fmt.Println("Masukkan koordinat titik sembarang:")
	fmt.Scan(&x234, &y234)
	titik234 := Titik234{x234, y234}

	dalamLingkaran1234 := diDalam234(lingkaran1234, titik234)
	dalamLingkaran2234 := diDalam234(lingkaran2234, titik234)

	if dalamLingkaran1234 && dalamLingkaran2234 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dalamLingkaran1234 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dalamLingkaran2234 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
