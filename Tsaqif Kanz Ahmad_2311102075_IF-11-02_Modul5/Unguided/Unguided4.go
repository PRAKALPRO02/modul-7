package main

import (
	"fmt"
	"strings"
)

const NMAX int = 127

type Tabel struct {
	tab [NMAX]rune
	m   int
}

func IsiArray(t *Tabel, n *int) {
	fmt.Println("Masukkan teks :")
	var input string
	fmt.Scanln(&input)

	for _, char := range input {
		if char == '.' || *n >= NMAX {
			break
		}
		if char != ' ' {
			t.tab[*n] = char
			*n++
		}
	}
	t.m = *n
}

func printArray(t Tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(string(t.tab[i]))
	}
	fmt.Println()
}

func reverseArray(t *Tabel, n int) {
	for i := 0; i < n/2; i++ {
		t.tab[i], t.tab[n-i-1] = t.tab[n-i-1], t.tab[i]
	}
}

func isPalindrome(t Tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if strings.ToLower(string(t.tab[i])) != strings.ToLower(string(t.tab[n-i-1])) {
			return false
		}
	}
	return true
}

func main() {
	var t Tabel
	var n int
	IsiArray(&t, &n)
	fmt.Print("Teks: ")
	printArray(t, n)
	fmt.Print("Reverse teks: ")
	reverseArray(&t, n)
	printArray(t, n)

	if isPalindrome(t, n) {
		fmt.Println("Palindrome: true")
	} else {
		fmt.Println("Palindrome: false")
	}
}
