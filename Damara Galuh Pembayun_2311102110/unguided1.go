package main

import (
        "fmt"
        "math"
)

// Titik represents a point in a 2D coordinate system
type Titik struct {
        x int
        y int
}

// Lingkaran represents a circle with a center and radius
type Lingkaran struct {
        pusat Titik
        radius int
}

// jarak calculates the distance between two points
func jarak(p, q Titik) float64 {
        return math.Sqrt(math.Pow(float64(p.x-q.x), 2) + math.Pow(float64(p.y-q.y), 2))
}

// didalam checks if a point is inside a circle
func didalam(c Lingkaran, p Titik) bool {
        return jarak(c.pusat, p) <= float64(c.radius)
}

func main() {
        var t1, t2, p Titik
        var r1, r2 int

        fmt.Scan(&t1.x, &t1.y, &r1)
        fmt.Scan(&t2.x, &t2.y, &r2)
        fmt.Scan(&p.x, &p.y)

        l1 := Lingkaran{pusat: t1, radius: r1}
        l2 := Lingkaran{pusat: t2, radius: r2}

        if didalam(l1, p) && didalam(l2, p) {
                fmt.Println("Titik di dalam lingkaran 1 dan 2")
        } else if didalam(l1, p) {
                fmt.Println("Titik di dalam lingkaran 1")
        } else if didalam(l2, p) {
                fmt.Println("Titik di dalam lingkaran 2")
        } else {
                fmt.Println("Titik di luar lingkaran 1 dan 2")
        }
}