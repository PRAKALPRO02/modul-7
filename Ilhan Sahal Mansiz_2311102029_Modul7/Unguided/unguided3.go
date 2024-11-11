package main

import (
    "fmt"
)

func main() {
    var klubA, klubB string
    var skorA, skorB int
    var pemenang []string
    var hasilPertandingan [][2]int

    fmt.Print("Masukkan nama klub A: ")
    fmt.Scan(&klubA)
    fmt.Print("Masukkan nama klub B: ")
    fmt.Scan(&klubB)

    fmt.Println("Masukkan skor untuk pertandingan, masukkan skor negatif untuk mengakhiri:")
    for {
        fmt.Scan(&skorA, &skorB)

        if skorA < 0 || skorB < 0 {
            fmt.Println("Pertandingan selesai")
            break
        }

        hasilPertandingan = append(hasilPertandingan, [2]int{skorA, skorB})
    }

    for _, hasil := range hasilPertandingan {
        skorA, skorB = hasil[0], hasil[1]

        if skorA > skorB {
            pemenang = append(pemenang, klubA)
        } else if skorB > skorA {
            pemenang = append(pemenang, klubB)
        } else {
            pemenang = append(pemenang, "Draw")
        }
    }

    fmt.Println("\nDaftar hasil pertandingan:")
    for i, hasil := range hasilPertandingan {
        fmt.Printf("Pertandingan %d: %s %d - %d %s\n", i+1, klubA, hasil[0], hasil[1], klubB)
    }

    fmt.Println("\nDaftar klub yang memenangkan pertandingan:")
    for i, klub := range pemenang {
        if klub == "Draw" {
            fmt.Printf("Pertandingan %d: Draw\n", i+1)
        } else {
            fmt.Printf("Pertandingan %d: %s\n", i+1, klub)
        }
    }
}