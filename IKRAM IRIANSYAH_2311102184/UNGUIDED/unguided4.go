package main

import (
	"fmt"
)

const NMAX int = 127
type tabel [NMAX]rune

func array() (t tabel, n int) {
	fmt.Println("Masukkan karakter (akhiri dengan titik) :")
	for n < NMAX {
		var char rune
		fmt.Printf("Karakter ke-%d : ", n+1)
		fmt.Scanf("%c\n", &char)
		if char == '.' {
			break
		}
		t[n] = char
		n++
	}
	return
}

func cetak_array(t tabel, n int, label string) {
	fmt.Printf("%s: ", label)
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t[i])
	}
	fmt.Println()
}

func balik_array(t tabel, n int) tabel {
	var reversed tabel
	for i := 0; i < n; i++ {
		reversed[i] = t[n-1-i]
	}
	return reversed
}

func palindrom(t tabel, n int) bool {
	reversed := balik_array(t, n)
	for i := 0; i < n; i++ {
		if t[i] != reversed[i] {
			return false
		}
	}
	return true
}

func main() {
	tab, m := array()
	cetak_array(tab, m, "Original")

	reversed := balik_array(tab, m)
	cetak_array(reversed, m, "Reversed")

	if palindrom(tab, m) {
		fmt.Println("Palindrom  True")
	} else {
		fmt.Println("Palindrom : False")
	}
}
