package main
import (
	"fmt"
	"sort"
)

// Struktur untuk menampung data mahasiswa
type mahasiswa struct {
				Nama string
				Matematika int	
				Fisika int
				Kimia int
				RataRata float64
}

// fungsi untuk menghitung rata - rata nilai tiap mahasiswa
func hitungRataRata(m *mahasiswa) {
					total := m.Matematika + m.Fisika + m.Kimia
					m.RataRata = float64(total) / 3.0
}

func main() {
				// array untuk menampung data mahasiswa
				mahasiswa := []mahasiswa{
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
									return mahasiswa[i].RataRata > mahasiswa[j].RataRata
				})

				// menampilkan hasil
				fmt.Println("Peringkat Mahasiswa berdasarkan rata-rata nilai:")
				for i, m := range mahasiswa {
										fmt.Printf("%d. %s - Rata-rata: %.2f (Matematika: %d, Fisika: %d, Kimia: %d)\n",
										i+1, m.Nama, m.RataRata, m.Matematika, m.Fisika, m.Kimia)
				}
}