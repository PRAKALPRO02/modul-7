package main

import (
	"fmt"
)

const NMAX int = 127

type tabel [NMAX]rune

func isi(t *tabel, n *int) {
	var ch rune
	fmt.Println("Masukkan karakter : ")

	for {
		_, err := fmt.Scanf("%c", &ch)
		if err != nil || ch == '.' || *n >= NMAX {
			break
		}
		t[*n] = ch
		*n++
	}
}
func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t[i])
	}
	fmt.Println()
}

func balik(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-i-1] = t[n-i-1], t[i]
	}
}

func palindrome(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-i-1] {
			return false
		}
	}
	return true
}

func main() {
	var tab213 tabel
	var n213 int = 0
	isi(&tab213, &n213)

	fmt.Print("Teks: ")
	cetakArray(tab213, n213)

	balik(&tab213, n213)
	fmt.Print("Reverse Teks: ")
	cetakArray(tab213, n213)

	if palindrome(tab213, n213) {
		fmt.Println("Palindrome: true")
	} else {
		fmt.Println("Palindrome: false")
	}
}
