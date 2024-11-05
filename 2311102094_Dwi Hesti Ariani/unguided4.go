package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

var tab tabel
var m int

func isiArray(t *tabel, n *int) {

	var i int
	var ch rune
	i = 0
	fmt.Scanf("%c", &ch)
	for ch != '.' && i < NMAX {
		t[i] = ch
		i++
		fmt.Scanf("%c", &ch)
	}
	*n = i
}

func cetakArray(t tabel, n int) {

	var i int
	for i = 0; i < n; i++ {
		fmt.Printf("%c", t[i])
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {

	var i, j int
	var temp rune
	for i = 0; i < n/2; i++ {
		j = n - i - 1
		temp = t[i]
		t[i] = t[j]
		t[j] = temp
	}
}

func isPalindrome(t tabel, n int) bool {

	for i := 0; i < n/2; i++ {
		if t[i] != t[n-i-1] {
			return false
		}
	}
	return true
}

func main() {
	isiArray(&tab, &m)
	fmt.Println("Array yang dimasukkan:")
	cetakArray(tab, m)

	if isPalindrome(tab, m) {
		fmt.Println("Array tersebut adalah palindrom.")
	} else {
		fmt.Println("Array tersebut bukan palindrom.")
	}
}
