package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var input string
	fmt.Print("Teks: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ = reader.ReadString('\n')

	if len(input) > 0 && input[len(input)-1] == '-' {
		input = input[:len(input)-1]
	}

	for _, ch := range input {
		if unicode.IsSpace(ch) {
			continue
		}
		if *n >= NMAX {
			break
		}
		t[*n] = ch
		*n++
	}
}

func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-1-i] = t[n-1-i], t[i]
	}
}

func palindrom(t tabel, n int) bool {
	var reversed tabel
	copy(reversed[:], t[:n])
	balikanArray(&reversed, n)
	for i := 0; i < n; i++ {
		if t[i] != reversed[i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	isiArray(&tab, &m)
	balikanArray(&tab, m)

	if palindrom(tab, m) {
		fmt.Println("Palindrom ? true")
	} else {
		fmt.Println("Palindrom ? false")
	}
}
