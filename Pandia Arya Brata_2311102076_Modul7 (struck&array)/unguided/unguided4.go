package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(tab *tabel, n *int) {
	fmt.Print("Masukkan teks (maksimal 127 karakter): ")
	input := bufio.NewReader(os.Stdin)
	text, _ := input.ReadString('\n')
	text = strings.TrimSpace(text)

	*n = len(text)
	for i := 0; i < *n && i < NMAX; i++ {
		tab[i] = rune(text[i])
	}
}

func cetakArray(tab tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c", tab[i])
	}
	fmt.Println()
}

func balikkanArray(tab *tabel, n int) {
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		tab[i], tab[j] = tab[j], tab[i]
	}
}

func palindrom(tab tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if tab[i] != tab[n-i-1] {
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

	balikkanArray(&tab, n)
	fmt.Print("Reverse: ")
	cetakArray(tab, n)

	if palindrom(tab, n) {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Bukan Palindrome")
	}
}
