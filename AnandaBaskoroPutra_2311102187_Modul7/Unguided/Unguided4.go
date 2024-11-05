package main

import "fmt"

const NMAX int = 127

type tabel struct {
	tab [NMAX]rune
	m   int
}

func isiArray(t *tabel, n *int) {
	
	fmt.Print("Masukkan karakter (akhiri dengan titik): ")
	var input rune
	for *n = 0; *n < NMAX; *n++ {
		fmt.Scanf("%c", &input)
		if input == '.' {
			break
		}
		t.tab[*n] = input
	}
	t.m = *n
}

func cetakArray(t tabel, n int) {
	
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t.tab[i])
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	
	for i := 0; i < n/2; i++ {
		t.tab[i], t.tab[n-i-1] = t.tab[n-i-1], t.tab[i]
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

	fmt.Print("Reverse teks: ")
	balikanArray(&tab, m)
	cetakArray(tab, m)

	isPalindrom := palindrom(tab, m)
	if isPalindrom {
		fmt.Println("Palindrom: true")
	} else {
		fmt.Println("Palindrom: false")
	}
}