package main

import "fmt"

const NMAX = 127

type tabel [NMAX]rune

func isiArray(t *tabel) int {
	var m int
	fmt.Print("Teks (tekan '.' untuk berhenti): ")
	for {
		var input string
		fmt.Scan(&input)
		if input == "." {
			break
		}
		t[m] = []rune(input)[0]
		m++
	}
	return m
}

func balikanArray(t *tabel, n int) {
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		t[i], t[j] = t[j], t[i]
	}
}

func cetakArray(t tabel, n int) {
	fmt.Print("Reverse teks: ")
	for i := 0; i < n; i++ {
		fmt.Print(string(t[i]), " ")
	}
	fmt.Println()
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
	m := isiArray(&tab)
	balikanArray(&tab, m)
	cetakArray(tab, m)

  if palindrom(tab, m) {
		fmt.Println("Palindrom? true")
	} else {
		fmt.Println("Palindrom? false")
	}
}
