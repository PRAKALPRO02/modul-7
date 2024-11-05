package main

import (
	"fmt"
)

const NMAX int = 127
type tabel [NMAX]rune

// Fungsi untuk mengisi array dengan karakter hingga titik ditemukan atau mencapai batas NMAX
func isiArray(t *tabel, n *int) {
	var input rune
	*n = 0

	fmt.Println("Masukkan karakter (akhiri dengan titik '.'): ")
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

// Fungsi untuk membalikkan isi array
func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-i-1] = t[n-i-1], t[i]
	}
}

// Fungsi untuk memeriksa apakah array merupakan palindrom
func isPalindrome(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-i-1] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	// Mengisi array tab dengan karakter dari pengguna
	isiArray(&tab, &m)

	// Mencetak isi array asli
	fmt.Print("Teks: ")
	cetakArray(tab, m)

	// Membalikkan isi array
	balikanArray(&tab, m)

	// Mencetak isi array yang sudah dibalik
	fmt.Print("Reverse teks: ")
	cetakArray(tab, m)

	// Memeriksa apakah array merupakan palindrom
	if isPalindrome(tab, m) {
		fmt.Println("Teks adalah palindrom.")
	} else {
		fmt.Println("Teks bukan palindrom.")
	}
}
