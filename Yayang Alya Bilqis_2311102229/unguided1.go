package main

import (
	"fmt"
	"math"
)

func jarak(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}

func diDalamLingkaran(px, py, cx, cy, r float64) bool {
	return jarak(px, py, cx, cy) <= r
}

func main() {
	var cx1, cy1, r1 float64
	var cx2, cy2, r2 float64
	var n int

	fmt.Print("Masukkan pusat dan radius lingkaran 1: ")
	fmt.Scan(&cx1, &cy1, &r1)

	fmt.Print("Masukkan pusat dan radius lingkaran 2: ")
	fmt.Scan(&cx2, &cy2, &r2)

	fmt.Print("Masukkkan jumlah titik: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var px, py float64
		fmt.Printf("Masukkan koordinat titik: ", i+1)
		fmt.Scan(&px, &py)

		diLingkaran1 := diDalamLingkaran(px, py, cx1, cy1, r1)
		diLingkaran2 := diDalamLingkaran(px, py, cx2, cy2, r2)

		if diLingkaran1 && diLingkaran2 {
			fmt.Println("Titik di dalam lingkaran 1 dan 2")
		} else if diLingkaran1 {
			fmt.Println("Titik di dalam lingkaran 1")
		} else if diLingkaran2 {
			fmt.Println("Titik di dalam lingkaran 2")
		} else {
			fmt.Println("Titik di luar lingkaran 1 dan 2")
		}
	}
}
