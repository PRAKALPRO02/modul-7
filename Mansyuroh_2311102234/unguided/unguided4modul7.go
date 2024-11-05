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
	var input string
	fmt.Print("Masukkan teks: ")
	fmt.Scanf("%s", &input)

	t.m = len(input)
	*n = t.m
	for i := 0; i < *n && i < NMAX; i++ {
		t.tab[i] = rune(input[i])
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t.tab[i])
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		j := n - i - 1
		t.tab[i], t.tab[j] = t.tab[j], t.tab[i]
	}
}

func palindrom(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t.tab[i] != t.tab[n-i-1] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	isiArray(&tab, &m)

	fmt.Print("Teks: ")
	cetakArray(tab, m)

	balikanArray(&tab, m)
	fmt.Print("Reverse teks: ")
	cetakArray(tab, m)

	isPalindrome := palindrom(tab, m)
	fmt.Print("Palindrome? ")
	fmt.Println(isPalindrome)
}
