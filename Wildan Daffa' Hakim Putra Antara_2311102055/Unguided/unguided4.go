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

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t[i])
		fmt.Print(" ")
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	var dat tabel

	for i := n - 1; i >= 0; i-- {
		dat[(n-1)-i] = t[i]
	}
	copy((*t)[:], dat[:])
}

func main() {
	var tab tabel
	var m int
	fmt.Print("Teks\t\t: ")
	isiArray(&tab, &m)
	balikanArray(&tab, m)
	fmt.Print("Reverse Teks\t: ")
	cetakArray(tab, m)
}
