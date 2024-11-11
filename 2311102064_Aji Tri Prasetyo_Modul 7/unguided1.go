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
func diDalam(c Lingkaran, p Titik) bool {
	return jarak(c.pusat, p) <= float64(c.radius)
}

func main() {
	var l1, l2 Lingkaran
	var t Titik
	fmt.Print("Masukkan koordinat pusat dan radius lingkaran 1 (cx1 cy1 r1): ")
	fmt.Scan(&l1.pusat.x, &l1.pusat.y, &l1.radius)

	fmt.Print("Masukkan koordinat pusat dan radius lingkaran 2 (cx2 cy2 r2): ")
	fmt.Scan(&l2.pusat.x, &l2.pusat.y, &l2.radius)

	fmt.Print("Masukkan koordinat titik sembarang (x y): ")
	fmt.Scan(&t.x, &t.y)
	diDalamL1 := diDalam(l1, t)
	diDalamL2 := diDalam(l2, t)

	if diDalamL1 && diDalamL2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if diDalamL1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if diDalamL2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
