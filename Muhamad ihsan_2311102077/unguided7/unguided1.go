package main

import (
	"fmt"
)

type Titik struct {
	x int
	y int
}

type lingkaran struct {
	center Titik
	radius int
}

func isTitikInlingkaran(p Titik, c lingkaran) bool {

	jarak := (p.x-c.center.x)*(p.x-c.center.x) + (p.y-c.center.y)*(p.y-c.center.y)

	radius := c.radius * c.radius
	return jarak < radius
}

func main() {
	var lingkaran1 lingkaran
	var lingkaran2 lingkaran
	var titik Titik 

	fmt.Scan(&lingkaran1.center.x, &lingkaran1.center.y, &lingkaran1.radius)
	fmt.Scan(&lingkaran2.center.x, &lingkaran2.center.y, &lingkaran2.radius)
	fmt.Scan(&titik.x, &titik.y)

	inlingkaran1 := isTitikInlingkaran(titik, lingkaran1)
	inlingkaran2 := isTitikInlingkaran(titik, lingkaran2)

	if inlingkaran1 && inlingkaran2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if inlingkaran1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if inlingkaran2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
