package main

import (
	"fmt"
)

const NMAX int = 127

type tabel [NMAX]rune

// Fungsi untuk mengisi array dengan karakter dari input user
func isiArray(t *tabel, n *int) {
	var ch rune
	*n = 0
	for {
		fmt.Scanf("%c", &ch)
		if ch == '\n' || ch == 'T' {
			break
		}
		t[*n] = ch
		*n++
		if *n >= NMAX {
			break
		}
	}
}

// Fungsi untuk membalikkan isi array
func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-1-i] = t[n-1-i], t[i]
	}
}

// Fungsi untuk mencetak isi array
func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t[i])
	}
	fmt.Println()
}

// Fungsi untuk memeriksa apakah array membentuk palindrom
func palindrom(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var n int

	// Mengisi array dengan input dari user
	fmt.Print("Masukkan teks: ")
	isiArray(&tab, &n)

	// Mencetak array asli
	fmt.Print("Teks: ")
	cetakArray(tab, n)

	// Membalikkan isi array
	balikanArray(&tab, n)
	fmt.Print("Reverse teks: ")
	cetakArray(tab, n)

	// Memeriksa apakah array membentuk palindrom
	if palindrom(tab, n) {
		fmt.Println("Teks ini adalah palindrom.")
	} else {
		fmt.Println("Teks ini bukan palindrom.")
	}
}
