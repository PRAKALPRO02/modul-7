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
func palindrom(t tabel, n int) bool {
	var reversed tabel
	copy(reversed[:], t[:n])
	balikanArray(&reversed, n)

	for i := 0; i < n; i++ {
		if t[i] != reversed[i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	fmt.Print("Teks: ")
	isiArray(&tab, &m)

	fmt.Print("Teks: ")
	cetakArray(tab, m)

	isPalindrome := palindrom(tab, m)
	fmt.Printf("palindrom? %t\n", isPalindrome)
}
