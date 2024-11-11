package main

import (
	"fmt"
	"sort"
)

type Mahasiswa struct {
	nama       string
	matematika int
	fisika     int
	kimia      int
	rataRata   float64
}

func hitungRataRata(m *Mahasiswa) {
	total := m.matematika + m.fisika + m.kimia
	m.rataRata = float64(total) / 3.0
}

func main() {
	mahasiswa := []Mahasiswa{
		{"Ali", 85, 90, 80, 0},
		{"Budi", 70, 75, 80, 0},
		{"Cici", 90, 85, 95, 0},
		{"Doni", 60, 65, 70, 0},
		{"Eka", 100, 95, 90, 0},
	}

	for i := range mahasiswa {
		hitungRataRata(&mahasiswa[i])
	}

	sort.Slice(mahasiswa, func(i, j int) bool {
		return mahasiswa[i].rataRata > mahasiswa[j].rataRata
	})

	fmt.Println("Peringkat mahasiswa berdasarkan nilai rata-rata: ")
	for i, m := range mahasiswa {
		fmt.Printf("%d. %s - rata-rata: %.2f(matematika: %d, fisika: %d, kimia: %d)\n", i+1, m.nama, m.rataRata, m.matematika, m.fisika, m.kimia)

	}
}
