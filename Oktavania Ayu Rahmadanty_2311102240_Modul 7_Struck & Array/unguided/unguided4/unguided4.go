package main

import (
	"fmt"
)

const NMAX int = 127

type tabel struct {
	tab [NMAX]rune
	m   int
}

func isiArray(t *tabel, n *int) {
	var input rune
	fmt.Print("Masukkan karakter (akhiri dengan '.'): ")
	for {
		fmt.Scanf("%c", &input)
		if input == '.' || *n >= NMAX {
			break
		}
		t.tab[*n] = input
		*n++
	}
	t.m = *n
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t.tab[i])
	}
	fmt.Println()
}

func balikkanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t.tab[i], t.tab[n-i-1] = t.tab[n-i-1], t.tab[i]
	}
}

func palindrome(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t.tab[i] != t.tab[n-i-1] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var n int

	isiArray(&tab, &n)

	fmt.Print("Teks: ")
	cetakArray(tab, n)

	fmt.Print("Reverse teks: ")
	balikkanArray(&tab, n)
	cetakArray(tab, n)

	if palindrome(tab, n) {
		fmt.Println("Palindrome: true")
	} else {
		fmt.Println("Palindrome: false")
	}
}
