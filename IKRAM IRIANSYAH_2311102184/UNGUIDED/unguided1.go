package main

import (
    "fmt"
    "math"
)

type Titik struct {
    x, y int
}

type Lingkaran struct {
    pusat Titik
    radius int
}

func jarak(p, q Titik) float64 {
    return math.Sqrt(float64((p.x - q.x)*(p.x - q.x) + (p.y - q.y) * (p.y - q.y)))
}

func inside(c Lingkaran, p Titik) bool {
    return jarak(c.pusat, p) <= float64(c.radius)
}

func main() {
    var l1, l2 Lingkaran
    var p Titik

    fmt.Scan(&l1.pusat.x, &l1.pusat.y, &l1.radius)
    fmt.Scan(&l2.pusat.x, &l2.pusat.y, &l2.radius)
    fmt.Scan(&p.x, &p.y)

    insideL1 := inside(l1, p)
    insideL2 := inside(l2, p)

    if insideL1 && insideL2 {
        fmt.Println("Titik di dalam lingkaran 1 dan 2")
    } else if insideL1 {
        fmt.Println("Titik di dalam lingkaran 1")
    } else if insideL2 {
        fmt.Println("Titik di dalam lingkaran 2")
    } else {
        fmt.Println("Titik di luar lingkaran 1 dan 2")
    }
}
