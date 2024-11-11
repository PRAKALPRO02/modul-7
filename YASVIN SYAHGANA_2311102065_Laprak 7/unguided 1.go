package main

import (
	"fmt"
	"math"
)

type titik struct {
	x, y int
}

type lingkaran struct {
	cx, cy, r int
}

func jarak(p titik, c lingkaran) float64 {

	return math.Sqrt(float64((p.x-c.cx)*(p.x-c.cx) + (p.y-c.cy)*(p.y-c.cy)))
}

func posisiTitik(p titik, l1, l2 lingkaran) string {
	jarak1 := jarak(p, l1)
	jarak2 := jarak(p, l2)

	dalam1 := jarak1 < float64(l1.r)
	dalam2 := jarak2 < float64(l2.r)

	if dalam1 && dalam2 {
		return "Titik di dalam lingkaran 1 dan 2"
	} else if dalam1 {
		return "Titik di dalam lingkaran 1"
	} else if dalam2 {
		return "Titik di dalam lingkaran 2"
	} else {
		return "Titik di luar lingkaran 1 dan 2"
	}
}

func main() {
	var l1, l2 lingkaran
	var p titik

	fmt.Print("Masukkan koordinat pusat dan radius lingkaran 1 (cx cy r): ")
	fmt.Scan(&l1.cx, &l1.cy, &l1.r)

	fmt.Print("Masukkan koordinat pusat dan radius lingkaran 2 (cx cy r): ")
	fmt.Scan(&l2.cx, &l2.cy, &l2.r)

	fmt.Print("Masukkan koordinat titik sembarang (x y): ")
	//fahrur059
	fmt.Scan(&p.x, &p.y)
	fmt.Print("\n")

	hasil := posisiTitik(p, l1, l2)
	fmt.Println(hasil)
}
