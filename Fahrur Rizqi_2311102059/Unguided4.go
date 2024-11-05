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
		if char == '.' || *n >= NMAX {
			break
		}
		t.tab[*n] = char
		*n++
	}
	t.m = *n
}

// Prosedur untuk mencetak isi array
func cetakArray(t tabel, n int) {
	fmt.Print("Reverse teks :")
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t.tab[i])
	}
	fmt.Println()
}

// Prosedur untuk membalik isi array
func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t.tab[i], t.tab[n-i-1] = t.tab[n-i-1], t.tab[i]
	}
}

func main() {
	var tab tabel
	var n int

	// Isi array tab dengan memanggil prosedur isiArray
	isiArray(&tab, &n)

	// Balikkan isi array tab dengan memanggil balikanArray
	balikanArray(&tab, n)

	// Cetak isi array tab
	cetakArray(tab, n)
}
