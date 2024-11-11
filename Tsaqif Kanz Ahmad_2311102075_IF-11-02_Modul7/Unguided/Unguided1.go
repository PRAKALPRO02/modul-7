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

func isInCircle(p Point, c Circle) bool {
	return distance(p, c.center) <= float64(c.radius)
}

func main() {
	var c1, c2 Circle
	var p Point

	// Read input
	fmt.Scan(&c1.center.x, &c1.center.y, &c1.radius)
	fmt.Scan(&c2.center.x, &c2.center.y, &c2.radius)
	fmt.Scan(&p.x, &p.y)

	// Determine the position of the point
	if isInCircle(p, c1) && isInCircle(p, c2) {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if isInCircle(p, c1) {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if isInCircle(p, c2) {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
