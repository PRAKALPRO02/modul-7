package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var input string
	fmt.Print("Masukkan teks (akhiri dengan '.'): ")
	fmt.Scanln(&input)
	*n = 0
	for _, ch := range input {
		if ch == '.' || *n >= NMAX {
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

func balikkanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-i-1] = t[n-i-1], t[i]
	}
}

func palindrom(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-i-1] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var n int
	isiArray(&tab, &n)
	fmt.Print("Teks yang dimasukkan: ")
	cetakArray(tab, n)

	if palindrom(tab, n) {
		fmt.Println("Apakah teks ini palindrom? Ya")
	} else {
		fmt.Println("Apakah teks ini palindrom? Tidak")
	}
}
