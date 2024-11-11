package main

import (
	"fmt"
)

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Scanf("%c", &t[i])
	}
}

func cetakArray(t *tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(string(t[i]))
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-i-1] = t[n-i-1], t[i]
	}
}

func isPalindrome(t *tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-i-1] {
			return false
		}
	}
	return true
}

func main() {
	var t tabel
	var n int

	fmt.Print("Masukkan jumlah karakter (maksimal 127): ")
	fmt.Scan(&n)

	fmt.Println("Masukkan karakter:")
	var dummy byte
	fmt.Scanf("%c", &dummy)

	isiArray(&t, n)

	fmt.Print("Array yang dimasukkan: ")
	cetakArray(&t, n)

	balikanArray(&t, n)
	fmt.Print("Array setelah dibalik: ")
	cetakArray(&t, n)

	if isPalindrome(&t, n) {
		fmt.Println("Array adalah palindrom")
	} else {
		fmt.Println("Array bukan palindrom")
	}
}