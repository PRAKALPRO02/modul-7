package main

import "fmt"

type Pertandingan struct {
    klub1, klub2 string
    hasil []string
}

func main() {
    var p Pertandingan
    var skorA_2311102181, skorB int

    fmt.Print("Masukkan nama Klub A: ")
    fmt.Scanln(&p.klub1)
    fmt.Print("Masukkan nama Klub B: ")
    fmt.Scanln(&p.klub2)

    // Memulai loop untuk input skor pertandingan
    for j := 1; ; j++ {
        fmt.Printf("Masukkan skor pertandingan %v : ", j)
        fmt.Scanln(&skorA_2311102181, &skorB)
        
        // Mengakhiri loop jika skor negatif ditemukan
        if skorA_2311102181 < 0 || skorB < 0 {
            break
        }

        // Menentukan hasil pertandingan berdasarkan skor
        switch {
        case skorA_2311102181 > skorB:
            p.hasil = append(p.hasil, p.klub1)
        case skorA_2311102181 < skorB:
            p.hasil = append(p.hasil, p.klub2)
        default:
            p.hasil = append(p.hasil, "Draw")
        }
    }

    // Menampilkan hasil setiap pertandingan
    for k, hasil := range p.hasil {
        fmt.Printf("Hasil pertandingan %d: %v\n", k+1, hasil)
    }

    fmt.Println("Semua pertandingan selesai.")
}


