package main

import (
	"fmt"
	"math"
)

func main() {
	var cx1, cy1, r1, cx2, cy2, r2, x, y int
	fmt.Scan(&cx1, &cy1, &r1, &cx2, &cy2, &r2, &x, &y)

	dalamLingkaran1 := titikdalamlingkaran(cx1, cy1, r1, x, y)
	dalamLingkaran2 := titikdalamlingkaran(cx2, cy2, r2, x, y)

	if dalamLingkaran1 && dalamLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dalamLingkaran1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dalamLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}

func hitungJarak_219(cx1, cy1, cx2, cy2 int) int {
	return int(math.Sqrt(float64((cx1-cx2)*(cx1-cx2) + (cy1-cy2)*(cy1-cy2))))
}

func titikdalamlingkaran(tengahX, tengahY, jariJari, x, y int) bool {
	return hitungJarak_219(tengahX, tengahY, x, y) < jariJari
}
