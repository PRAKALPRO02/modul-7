package main

import (
	"fmt"
	"math"
)

type titik struct {
	x int
	y int
}

type lingkaran struct {
	pusat  titik
	radius int
}

func jarak(p, q titik) float64 {
	return math.Pow(float64(p.x-q.x), 2) + math.Pow(float64(p.y-q.y), 2)
}

func didalam(c lingkaran, p titik) bool {
	var jarak_titik = math.Sqrt(jarak(c.pusat, p))
	if jarak_titik < float64(c.radius) {
		return true
	}
	return false
}

func main() {
	var lingkaran1, lingkaran2 lingkaran
	var titikSembarang titik
	var hasil1, hasil2 bool
	fmt.Print("Titik pusat (x,y) dan jejari lingkaran 1 : ")
	fmt.Scanln(&lingkaran1.pusat.x, &lingkaran1.pusat.y, &lingkaran1.radius)
	fmt.Print("Titik pusat (x,y) dan jejari lingkaran 2 : ")
	fmt.Scanln(&lingkaran2.pusat.x, &lingkaran2.pusat.y, &lingkaran2.radius)
	fmt.Print("Titik sembarang (x,y) : ")
	fmt.Scanln(&titikSembarang.x, titikSembarang.y)
	hasil1 = didalam(lingkaran1, titikSembarang)
	hasil2 = didalam(lingkaran2, titikSembarang)
	if hasil1 && hasil2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if hasil1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if hasil2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
