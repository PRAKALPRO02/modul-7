package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	fmt.Print("Jumlah karakter [k < 127]: ")
	fmt.Scanln(n)
	if *n > NMAX {
		fmt.Println("Jumlah karakter terlalu banyak!")
	} else {
		fmt.Print("Teks: ")
		for i := 0; i < *n; i++ {
			fmt.Scanf("%c", &t[i])
		}
	}
}

func balikanArray(t *tabel, n int) {
	var temp tabel
	for i := 0; i < n; i++ {
		temp[i] = (*t)[n-1-i]
	}
	*t = temp
}

func palindrom(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-1-i] {
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
		fmt.Println("Palindrom: true")
	} else {
		fmt.Println("Palindrom: false")
	}
}
