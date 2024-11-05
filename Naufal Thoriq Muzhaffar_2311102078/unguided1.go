package main

import (
	"fmt"
	"math"
)

func main() {
	var ling1 Lingkaran
	var ling2 Lingkaran
	var ttk Titik
	fmt.Scan(&ling1.cx, &ling1.cy, &ling1.r)
	fmt.Scan(&ling2.cx, &ling2.cy, &ling2.r)
	fmt.Scan(&ttk.x, &ttk.y)

	lingkaran1 := ling1
	lingkaran2 := ling2
	titik := ttk

	hasil := cekTitik(lingkaran1, lingkaran2, titik)
	fmt.Println(hasil)
}

type Lingkaran struct {
	cx, cy, r float64
}

type Titik struct {
	x, y float64
}

func jarak(p, q Titik) float64 {
	return math.Sqrt((p.x-q.x)*(p.x-q.x) + (p.y-q.y)*(p.y-q.y))
}

func didalam(c Lingkaran, p Titik) bool {
	return jarak(Titik{c.cx, c.cy}, p) <= c.r
}

func cekTitik(c1, c2 Lingkaran, p Titik) string {
	inCircle1 := didalam(c1, p)
	inCircle2 := didalam(c2, p)

	if inCircle1 && inCircle2 {
		return "Titik di dalam lingkaran 1 dan 2"
	} else if inCircle1 {
		return "Titik di dalam lingkaran 1"
	} else if inCircle2 {
		return "Titik di dalam lingkaran 2"
	} else {
		return "Titik di luar lingkaran 1 dan 2"
	}
}
