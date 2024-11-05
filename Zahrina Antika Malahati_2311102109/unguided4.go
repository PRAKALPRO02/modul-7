package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	// Implementasi untuk mengisi array t dengan input dari pengguna
	fmt.Print("Masukkan teks: ")
	var input string
	fmt.Scanln(&input)
	*n = len(input)
	for i := 0; i < *n; i++ {
		(*t)[i] = rune(input[i])
	}
}

func cetakArray(t tabel, n int) {
	// Implementasi untuk mencetak isi array t
	fmt.Print("Teks: ")
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t[i])
	}
	fmt.Println()
}

func balikArray(t *tabel, n int) {
	// Implementasi untuk membalikkan isi array t
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		t[i], t[j] = t[j], t[i]
	}
}

func palindrom(t tabel, n int) bool {
	// Implementasi untuk mengecek apakah t adalah palindrom
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-i-1] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	isiArray(&tab, &m)
	fmt.Println("Reverse teks: ")
	balikArray(&tab, m)
	cetakArray(tab, m)

	fmt.Println("Palindrom:", palindrom(tab, m))
}

// Zahrina Antika Malahati_2311102109
