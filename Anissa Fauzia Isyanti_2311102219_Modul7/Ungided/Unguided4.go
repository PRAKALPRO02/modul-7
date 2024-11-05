package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

// Fungsi untuk mengisi array dengan karakter hingga titik ditemukan atau mencapai batas NMAX
func isiArray(t *tabel, n *int) {
	var input rune
	*n = 0

	// Meminta input dari pengguna
	for *n < NMAX {
		fmt.Scanf("%c", &input)
		if input == '.' {
			break
		}
		t[*n] = input
		*n++
	}
}

// Fungsi untuk mencetak isi array
func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t[i])
	}
	fmt.Println()
}

// Fungsi untuk membalikkan urutan isi array
func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-i-1] = t[n-i-1], t[i]
	}
}

func main() {
	var tab tabel
	var m int

	fmt.Print("Teks: ")
	// Mengisi array tab dengan memanggil fungsi isiArray
	isiArray(&tab, &m)

	// Mencetak teks asli
	fmt.Print("Teks: ")
	cetakArray(tab, m)

	// Membalikkan isi array tab
	balikanArray(&tab, m)

	// Mencetak hasil yang dibalik
	fmt.Print("Reverse teks: ")
	cetakArray(tab, m)
}
