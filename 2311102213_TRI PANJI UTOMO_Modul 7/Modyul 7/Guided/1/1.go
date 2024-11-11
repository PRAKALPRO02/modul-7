package main

import (
        "fmt"
        "sort"
)

// Struktur untuk menampung data mahasiswa
type Mahasiswa struct {
        Nama     string
        Matematika int
        Fisika    int
        Kimia     int
        Ratarata  float64
}

// Fungsi untuk menghitung rata-rata nilai tiap mahasiswa
func hitungRataRata(m *Mahasiswa) {
        total := m.Matematika + m.Fisika + m.Kimia
        m.Ratarata = float64(total) / 3.0
}

func main() {
        mahasiswa := []Mahasiswa{
			{"Ali", 85, 90, 80, 0},
			{"Budi", 70, 75, 80,  0},
			{"Cici", 90, 85, 95, 0},
			{"Doni", 60, 65, 70, 0},
			{"Eka", 100, 95, 90, 0},
		}

		for i := range mahasiswa {
			hitungRataRata(&mahasiswa[i])
		}

		// Mengurutkan mahasiswa berdasarkan nilai rata-rata (descending)
sort.Slice(mahasiswa, func(i, j int) bool {
    return mahasiswa[i].Ratarata > mahasiswa[j].Ratarata
})

// Menampilkan hasil
fmt.Println("Peringkat mahasiswa berdasarkan rata-rata nilai:")
for i, m := range mahasiswa {
	fmt.Printf("%d. %s - Rata-rata: %.2f (Matematika: %d, Fisika %d, Kimia %d)\n",
	i+1, m.Nama, m.Ratarata, m.Matematika, m.Fisika, m.Kimia)
	}
    

}