package main

import (
	"fmt"
)

const NMAX int = 127

type Tabel [NMAX]rune

// Fungsi untuk mengisi array dengan karakter dari input user
func isiArray(t *Tabel, n *int) {
	var input string
	fmt.Print("Teks: ")
	fmt.Scanln(&input)

	*n = len(input)
	for i := 0; i < *n && i < NMAX; i++ {
		t[i] = rune(input[i])
	}
}

// Fungsi untuk mencetak array karakter
func cetakArray(t Tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t[i])
	}
	fmt.Println()
}

// Fungsi untuk membalikkan urutan karakter dalam array
func balikkanArray(t *Tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-i-1] = t[n-i-1], t[i]
	}
}

func main() {
	var tab Tabel
	var n int

	// Mengisi array dengan karakter dari input
	isiArray(&tab, &n)

	// Menampilkan array asli
	fmt.Print("Teks: ")
	cetakArray(tab, n)

	// Membalikkan array dan menampilkannya
	balikkanArray(&tab, n)
	fmt.Print("Reverse teks: ")
	cetakArray(tab, n)
}
