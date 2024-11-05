package main

import (
	"fmt"
	"sort"
)

// struktur untuk menampung data mahasiswa
type Mahasiswa struct {
	Nama       string
	Matematika int
	Fisika     int
	Kimia      int
	RataRata   float64
}

// Fungsi untuk menghitung rata-rata nilai tiap mahasiswa
func hitungRataRata(m *Mahasiswa) {
	total := m.Matematika + m.Fisika + m.Kimia
	m.RataRata = float64(total) / 3.0
}

// Fungsi utama untuk mengelola dan mengurutkan data mahasiswa
func main() {
	//array untuk menampung data mahasiswa
	mahasiswa := []Mahasiswa{
		{"Ali", 85, 90, 80, 0},
		{"Budi", 70, 75, 80, 0},
		{"Cici", 90, 85, 95, 0},
		{"Doni", 60, 65, 70, 0},
		{"Eka", 100, 95, 90, 0},
	}
	//Menghitung rata-rata nilai tiap mahasiswa
	for i := range mahasiswa {
		hitungRataRata(&mahasiswa[i])
	}

	//Mengurutkan mahasiswa berdasarkan nilai rata-rata (descend)
	sort.Slice(mahasiswa, func(i, j int) bool {
		return mahasiswa[i].RataRata > mahasiswa[j].RataRata
	})

	//Menampilkan hasil
	fmt.Println("Peringkat mahasiswa berdasarkan rata-rata nilai: ")
	for i, m := range mahasiswa {
		fmt.Printf("%d. %s - Rata-rata: %.2f (Matematika: %d, Fisika: %d, Kimia: %d)\n",
			i+1, m.Nama, m.RataRata, m.Matematika, m.Fisika, m.Kimia)

	}
}
