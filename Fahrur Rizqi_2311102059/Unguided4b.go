package main

import (
	"fmt"
)

const NMAX int = 127

type tabel struct {
	tab [NMAX]rune
	m   int
}

// Prosedur untuk mengisi array
func isiArray(t *tabel, n *int) {
	var char rune
	*n = 0
	fmt.Print("Teks         : ")
	for {
		fmt.Scanf("%c", &char)
		if char == '\n' || *n >= NMAX { // Ganti '.' dengan '\n' untuk mengakhiri input dengan Enter
			break
		}
		t.tab[*n] = char
		*n++
	}
	t.m = *n
}

// Fungsi untuk mengecek apakah teks adalah palindrom
func palindrom(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t.tab[i] != t.tab[n-i-1] {
			return false
		}
	}
	return true
}

// Prosedur untuk mencetak isi array
func cetakArray(t tabel, n int) {
	fmt.Print("Reverse teks : ")
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t.tab[i])
	}
	fmt.Println()
}

func main() {
	var tab tabel
	var n int

	// Isi array tab dengan memanggil prosedur isiArray
	isiArray(&tab, &n)

	// Cek apakah array tab adalah palindrom
	isPalindrom := palindrom(tab, n)

	// Cetak hasil palindrom
	fmt.Print("Palindrom    : ")
	if isPalindrom {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
