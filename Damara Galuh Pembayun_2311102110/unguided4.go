package main

import "fmt"

const NMAX int = 127
type tabel [NMAX] rune

func isiArray(t *tabel, n *int) {
        // ... (asumsikan sudah diimplementasikan)
}

func cetakArray(t tabel, n int) {
        // ... (asumsikan sudah diimplementasikan)
}

func balikanArray(t tabel, n int) {
        // Membalikkan isi array t
        for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
                t[i], t[j] = t[j], t[i]
        }
}

func palindrom(t tabel, n int) bool {
        // Memeriksa apakah t adalah palindrom
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

        // Isi array tab dengan memanggil prosedur isiArray
        isiArray(&tab, &m)

        // Balikan isi array tab dengan memanggil balikanArray
        balikanArray(tab, m)

        // Cetak isi array tab
        cetakArray(tab, m)

        // Periksa apakah tab adalah palindrom
        if palindrom(tab, m) {
                fmt.Println("Palindrom")
        } else {
                fmt.Println("Bukan Palindrom")
        }
}