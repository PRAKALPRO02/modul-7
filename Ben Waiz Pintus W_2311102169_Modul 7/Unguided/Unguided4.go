// Ben Waiz Pintus W
// 2311102169
package main

import "fmt"

const NMAX int = 127

type tabel struct {
	tab [NMAX]rune
	m   int
}

func isiArray(t *tabel, n *int) {
	/* I.S. Data tersedia dalam piranti masukan
	   F.S. Array t berisi sejumlah n karakter yang dimasukkan user,
	   Proses input selama karakter bukanlah TITIK dan n <= NMAX */
	fmt.Print("Masukkan karakter (akhiri dengan titik): ")
	var input rune
	for *n = 0; *n < NMAX; *n++ {
		fmt.Scanf("%c", &input)
		if input == '.' {
			break
		}
		t.tab[*n] = input
	}
	t.m = *n
}

func cetakArray(t tabel, n int) {
	/* I.S. Terdefinisi array t yang berisi sejumlah n karakter
	   F.S. n karakter dalam array muncul di layar */
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t.tab[i])
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	/* I.S. Terdefinisi array t yang berisi sejumlah n karakter
	   F.S. Urutan isi array t terbalik */
	for i := 0; i < n/2; i++ {
		t.tab[i], t.tab[n-i-1] = t.tab[n-i-1], t.tab[i]
	}
}

func palindrom(t tabel, n int) bool {
	/* Mengembalikan true apabila susunan karakter di dalam t membentuk palindrom,
	   dan false apabila sebaliknya. Petunjuk: Manfaatkan prosedur balikanArray */
	for i := 0; i < n/2; i++ {
		if t.tab[i] != t.tab[n-i-1] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	// Isi array tab dengan memanggil prosedur isiArray
	isiArray(&tab, &m)

	// Cetak isi array tab
	fmt.Print("Teks: ")
	cetakArray(tab, m)

	// Balikan isi array tab dengan memanggil balikanArray
	fmt.Print("Reverse teks: ")
	balikanArray(&tab, m)
	cetakArray(tab, m)

	// Cek apakah teks merupakan palindrom
	isPalindrom := palindrom(tab, m)
	if isPalindrom {
		fmt.Println("Palindrom: true")
	} else {
		fmt.Println("Palindrom: false")
	}
}
