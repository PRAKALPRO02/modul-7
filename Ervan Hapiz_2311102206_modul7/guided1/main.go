package main

import (
	"fmt"
	"sort"
)

type mahasiswa struct {
	nama       string
	matematika int
	fisika     int
	kimia      int
	rata_rata  float64
}

func HitungRata(n *mahasiswa) {
	n.rata_rata = float64(n.matematika+n.fisika+n.kimia) / float64(3)
}
func main() {
	mahasiswa := []mahasiswa{
		{"Ali", 85, 90, 80, 0},
		{"Budi", 70, 75, 80, 0},
		{"Cici", 90, 85, 95, 0},
		{"Doni", 60, 65, 70, 0},
		{"Eka", 100, 95, 90, 0},
	}

	for i := range mahasiswa {
		HitungRata(&mahasiswa[i])
	}

	sort.Slice(mahasiswa, func(i, j int) bool {
		return mahasiswa[i].rata_rata > mahasiswa[j].rata_rata
	})

	for i := range mahasiswa {
		fmt.Printf("%d. %s - Rata-rata : %.2f (Matematika : %d, fisika : %d, Kimia : %d\n", i+1, mahasiswa[i].nama, mahasiswa[i].rata_rata, mahasiswa[i].matematika, mahasiswa[i].fisika, mahasiswa[i].kimia)
	}
}
