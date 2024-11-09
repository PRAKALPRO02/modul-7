package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	fmt.Print("Jumlah karakter [k < 127]: ")
	fmt.Scanln(n)
	if *n > NMAX {
		fmt.Println("Jumlah karakter terlalu banyak!")
		isiArray(t, n)
	} else {
		fmt.Print("Teks: ")
		for i := 0; i < *n; i++ {
			fmt.Scanf("%c", &t[i])
		}
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c ", t[i])
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	var temp tabel
	for i := 0; i < n; i++ {
		temp[i] = (*t)[n-1-i]
	}
	*t = temp
}

func main() {
	var tab tabel
	var m int
	isiArray(&tab, &m)
	balikanArray(&tab, m)
	fmt.Print("Reverse teks: ")
	cetakArray(tab, m)
}
