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
    return math.Sqrt(math.Pow(float64(p.x-q.x), 2) + math.Pow(float64(p.y-q.y), 2))
}


func didalam(c Lingkaran, p Titik) bool {
    return jarak(c.pusat, p) <= float64(c.radius)
}

func main() {
    var cx1, cy1, r1, cx2, cy2, r2, px, py int

    fmt.Scan(&cx1, &cy1, &r1, &cx2, &cy2, &r2, &px, &py)

    c1 := Lingkaran{Titik{cx1, cy1}, r1}
    c2 := Lingkaran{Titik{cx2, cy2}, r2}
    p := Titik{px, py}

    inC1 := didalam(c1, p)
    inC2 := didalam(c2, p)

    if inC1 && inC2 {
        fmt.Println("Titik berada di dalam lingkaran 1 dan 2")
    } else if inC1 {
        fmt.Println("Titik berada di dalam lingkaran 1")
    } else if inC2 {
        fmt.Println("Titik berada di dalam lingkaran 2")
    } else {
        fmt.Println("Titik berada di luar lingkaran 1 dan 2")
    }
}