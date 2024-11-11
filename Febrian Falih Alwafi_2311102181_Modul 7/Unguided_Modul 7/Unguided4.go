package main

import "fmt"

const MaxLength int = 127

type Tabel struct {
    data [MaxLength]rune
    length_2311102181 int
}

// Fungsi untuk menerima input teks dan mengisinya ke dalam array hingga titik atau batas panjang
func inputTeks(t *Tabel) {
    fmt.Println("Masukkan teks (akhiri dengan TITIK):")
    var input rune
    for i := 0; ; i++ {
        fmt.Scanf("%c", &input)
        if input == '.' || i >= MaxLength {
            break
        }
        if input != ' ' {
            t.data[i] = input
            t.length_2311102181++
        }
    }
}

// Fungsi untuk mencetak seluruh teks yang ada dalam array
func tampilkanTeks(t Tabel) {
    for i := 0; i < t.length_2311102181; i++ {
        fmt.Print(string(t.data[i]))
    }
    fmt.Println()
}

// Fungsi untuk membalikkan teks dalam array
func balikTeks(t *Tabel) {
    for i := 0; i < t.length_2311102181/2; i++ {
        t.data[i], t.data[t.length_2311102181-i-1] = t.data[t.length_2311102181-i-1], t.data[i]
    }
}

// Fungsi untuk memeriksa apakah teks adalah palindrome
func cekPalindrome(t Tabel) bool {
    for i := 0; i < t.length_2311102181/2; i++ {
        if t.data[i] != t.data[t.length_2311102181-i-1] {
            return false
        }
    }
    return true
}

func main() {
    var t Tabel

    // Memasukkan dan mengisi array dengan teks
    inputTeks(&t)
    
    // Menampilkan teks asli
    fmt.Print("Teks: ")
    tampilkanTeks(t)
    
    // Menampilkan teks setelah dibalik
    fmt.Print("Reverse teks: ")
    balikTeks(&t)
    tampilkanTeks(t)

    // Mengecek apakah teks adalah palindrome
    if cekPalindrome(t) {
        fmt.Println("Palindrome: true")
    } else {
        fmt.Println("Palindrome: false")
    }
}


