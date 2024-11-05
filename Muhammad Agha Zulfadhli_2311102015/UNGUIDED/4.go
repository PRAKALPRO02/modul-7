package main

import (
	"fmt"
)

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var input rune
	*n = 0
	for *n < NMAX {
		fmt.Scanf("%c\n", &input)
		if input == '.' {
			break
		}
		t[*n] = input
		*n++
	}
}

func cetakArray(t tabel, n int) {
	fmt.Print("Teks\t\t: ")
	for i := 0; i < n; i++ {
		fmt.Printf("%c ", t[i])
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-i-1] = t[n-i-1], t[i]
	}
}

func palindrom(t tabel, n int) bool {
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

	// balikanArray(&tab, m)
	// fmt.Println("Isi array setelah dibalik:")
	// cetakArray(tab, m)
	isiArray(&tab, &m)
	cetakArray(tab, m)
	if palindrom(tab, m) {
		fmt.Println("Palindrom\t? true")
	} else {
		fmt.Println("Palindrom\t? false")
	}
}
