package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var temp_inp rune
	for *n < NMAX {
		fmt.Scanf("%c", &temp_inp)
		if temp_inp == '.' {
			break
		} else if temp_inp == ' ' {
			continue
		}
		t[*n] = temp_inp
		*n++
	}
}

func checkPolindrom(t tabel, n int) bool {
	for i := 0; i < n; i++ {
		if t[i] != t[(n-1)-i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int
	fmt.Print("Teks\t\t: ")
	isiArray(&tab, &m)
	fmt.Print("Palindrom\t? ", checkPolindrom(tab, m))
}
