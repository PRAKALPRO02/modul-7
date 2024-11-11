// MUHAMMAD RAGIEL PRASTYO
// 2311102183
package main
import (
	"fmt"
	"math"
)

type Point struct {
	x, y int
}

type Circle struct {
	center Point
	radius int
}

func distance(p1, p2 Point) float64 {
	return math.Sqrt(float64((p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y)))
}

func isInside(circle Circle, point Point) bool {
	return distance(circle.center, point) <= float64(circle.radius)
}

func checkPointPosition(circle1, circle2 Circle, point Point) string {
	inCircle1 := isInside(circle1, point)
	inCircle2 := isInside(circle2, point)

	switch {
	case inCircle1 && inCircle2:
		return "Titik di dalam lingkaran 1 dan 2"
	case inCircle1:
		return "Titik di dalam lingkaran 1"
	case inCircle2:
		return "Titik di dalam lingkaran 2"
	default:
		return "Titik di luar lingkaran 1 dan 2"
	}
}

func main() {
	var cx1, cy1, r1 int
	var cx2, cy2, r2 int
	var px, py int

	fmt.Print("Masukkan koordinat pusat dan radius lingkaran 1 (cx1 cy1 r1): ")
	fmt.Scan(&cx1, &cy1, &r1)

	fmt.Print("Masukkan koordinat pusat dan radius lingkaran 2 (cx2 cy2 r2): ")
	fmt.Scan(&cx2, &cy2, &r2)

	fmt.Print("Masukkan koordinat titik sembarang (px py): ")
	fmt.Scan(&px, &py)

	circle1 := Circle{center: Point{x: cx1, y: cy1}, radius: r1}
	circle2 := Circle{center: Point{x: cx2, y: cy2}, radius: r2}
	point := Point{x: px, y: py}

	result := checkPointPosition(circle1, circle2, point)
	fmt.Println(result)
}