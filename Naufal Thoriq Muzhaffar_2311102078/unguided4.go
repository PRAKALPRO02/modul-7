package main

import "fmt"

const NMAX int = 127

type tabel struct {
	tab [NMAX]rune
	m   int
}

func isiArray(t *tabel, n *int) {
	fmt.Println("Masukkan teks (akhiri dengan TITIK):")
	for {
		var input rune
		fmt.Scanf("%c", &input)
		if input == '.' || *n >= NMAX {
			break
		}

		if input != ' ' {
			t.tab[*n] = input
			*n++
		}
	}
	t.m = *n
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(string(t.tab[i]))
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
	var t tabel
	var n int
	isiArray(&t, &n)
	fmt.Print("Teks: ")
	cetakArray(t, n)
	fmt.Print("Reverse teks: ")
	balikkanArray(&t, n)
	cetakArray(t, n)

	if palindrome(t, n) {
		fmt.Println("Palindrome: true")
	} else {
		fmt.Println("Palindrome: false")
	}
}
