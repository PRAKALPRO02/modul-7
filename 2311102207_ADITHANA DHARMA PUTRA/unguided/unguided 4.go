package main
import "fmt"

const NMAX207 int = 127
type tabel [NMAX207]rune

func isiArray(t *tabel, n *int) {
	var karakter rune
	*n = 0
	fmt.Print("Teks\t\t: ")

	for *n < NMAX207 {
		fmt.Scanf("%c", &karakter)
		if karakter == '.' {
			break
		}

		if karakter != '\n' {
			t[*n] = karakter
			*n++
		}
	}
}

func cetakArray(t tabel, n int) {
	fmt.Print("Reverse teks\t:")
	for i := 0; i < n; i++ {
		fmt.Printf(" %c", t[i])
	}
	fmt.Println()
}

func balikanArray(t tabel, n int) tabel {
	var hasil tabel
	for i := 0; i < n; i++ {
		hasil[i] = t[n-1-i]
	}
	return hasil
}

func palindrom(t tabel, n int) string {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-1-i] {
			return "false"
		}
	}
	return "true"
}



func main() {
	var tab tabel
	var m int

	isiArray(&tab, &m)
	balikanArray(tab, m)
	cetakArray(tab, m)

	fmt.Printf("\n!MODIFIKASI!\n\n")
  	cek := palindrom(tab, m)
	fmt.Printf("palindrom\t?%s",cek)
}
