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

func isInCircle(c Circle, p Point) bool {
	return distance(c.center, p) <= float64(c.radius)
}

func checkPointPosition(c1, c2 Circle, p Point) string {
	inCircle1 := isInCircle(c1, p)
	inCircle2 := isInCircle(c2, p)

	if inCircle1 && inCircle2 {
		return "Titik berada di dalam lingkaran 1 dan lingkaran 2"
	} else if inCircle1 {
		return "Titik berada di dalam lingkaran 1"
	} else if inCircle2 {
		return "Titik berada di dalam lingkaran 2"
	} else {
		return "Titik berada di luar lingkaran 1 dan lingkaran 2"
	}
}

func main() {
	var circle1, circle2 Circle
	var point Point

	fmt.Println("Masukkan koordinat pusat (x, y) dan radius untuk lingkaran 1:")
	fmt.Scan(&circle1.center.x, &circle1.center.y, &circle1.radius)

	fmt.Println("Masukkan koordinat pusat (x, y) dan radius untuk lingkaran 2:")
	fmt.Scan(&circle2.center.x, &circle2.center.y, &circle2.radius)

	fmt.Println("Masukkan koordinat titik (x, y):")
	fmt.Scan(&point.x, &point.y)

	result := checkPointPosition(circle1, circle2, point)
	fmt.Println(result)
}
