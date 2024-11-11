package main

import (
	"fmt"
	"math"
)

type Titik struct {
	x, y int
}
type Lingkaran struct {
	pusat  Titik
	radius int
}

func jarak(p, q Titik) float64 {
	return math.Sqrt(float64((p.x-q.x)*(p.x-q.x) + (p.y-q.y)*(p.y-q.y)))
}
func didalamLingkaran(l Lingkaran, t Titik) bool {
	return jarak(l.pusat, t) <= float64(l.radius)
}
func posisiTitik(l1, l2 Lingkaran, t Titik) string {
	dalamL1 := didalamLingkaran(l1, t)
	dalamL2 := didalamLingkaran(l2, t)
	if dalamL1 && dalamL2 {
		return "Titik di dalam lingkaran 1 dan 2"
	} else if dalamL1 {
		return "Titik di dalam lingkaran 1"
	} else if dalamL2 {
		return "Titik di dalam lingkaran 2"
	} else {
		return "Titik di luar lingkaran 1 dan 2"
	}
}
func main() {
	var l1, l2 Lingkaran
	var t Titik
	fmt.Println("Masukkan koordinat pusat dan radius lingkaran 1:")
	fmt.Scan(&l1.pusat.x, &l1.pusat.y, &l1.radius)
	fmt.Println("Masukkan koordinat pusat dan radius lingkaran 2:")
	fmt.Scan(&l2.pusat.x, &l2.pusat.y, &l2.radius)
	fmt.Println("Masukkan koordinat titik sembarang:")
	fmt.Scan(&t.x, &t.y)

	hasil := posisiTitik(l1, l2, t)
	fmt.Println(hasil)
}
