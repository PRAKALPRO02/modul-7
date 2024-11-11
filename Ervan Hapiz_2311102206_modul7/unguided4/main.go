package main

import (
	"fmt"
)

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var char rune
	*n = 0
	fmt.Print("Masukkan karakter (gunakan (.) untuk berhenti): ")

	for *n < NMAX {
		fmt.Scanf("%c", &char)
		if char == '.' {
			break
		}

		if char != '\n' {
			t[*n] = char
			*n++
		}
	}
}

func balikanArray(t tabel, n int) tabel {
	var hasil tabel
	for i := 0; i < n; i++ {
		hasil[i] = t[n-1-i]
	}
	return hasil
}

func palindrom(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-1-i] {
			return false
		}
	}
	return true
}

func cetakArray(t tabel, n int) {
	var balik tabel = balikanArray(t, n)
	isPalindrome := palindrom(t, n)

	fmt.Print("Teks\t\t:")
	for i := 0; i < n; i++ {
		fmt.Printf(" %c", t[i])
	}
	fmt.Println()

	fmt.Print("Reverse teks\t:")
	for i := 0; i < n; i++ {
		fmt.Printf(" %c", balik[i])
	}
	fmt.Println()

	fmt.Print("Palindrom\t? ")
	if isPalindrome {
		fmt.Print("true")
	} else {
		fmt.Print("false")
	}
}

func main() {
	var tab tabel
	var m int

	isiArray(&tab, &m)
	cetakArray(tab, m)

}
