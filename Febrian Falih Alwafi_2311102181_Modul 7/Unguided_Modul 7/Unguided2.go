package main

import (
    "fmt"
    "math"
)

// Fungsi untuk menghapus elemen pada indeks tertentu dalam slice
func hapusItem(slice []int, pos int) []int {
    return append(slice[:pos], slice[pos+1:]...)
}

// Fungsi untuk menghitung rata-rata dari elemen-elemen dalam slice
func hitungRatarata_2311102181(slice []int) float64 {
    total := 0
    for _, nilai := range slice {
        total += nilai
    }
    return float64(total) / float64(len(slice))
}

// Fungsi untuk menghitung standar deviasi dari elemen-elemen dalam slice
func hitungStandarDeviasi(slice []int) float64 {
    mean := hitungRatarata_2311102181(slice)
    var varian float64
    for _, nilai := range slice {
        varian += math.Pow(float64(nilai)-mean, 2)
    }
    return math.Sqrt(varian / float64(len(slice)))
}

// Fungsi untuk menghitung jumlah kemunculan elemen tertentu dalam slice
func hitungFrekuensi(slice []int, target int) int {
    jumlah := 0
    for _, nilai := range slice {
        if nilai == target {
            jumlah++
        }
    }
    return jumlah
}

func main() {
    // Menerima jumlah elemen dari pengguna dan mengisi slice
    var jumlahElemen int
    fmt.Print("Masukkan jumlah elemen dalam array: ")
    fmt.Scan(&jumlahElemen)

    data := make([]int, jumlahElemen)
    fmt.Println("Masukkan elemen-elemen array:")
    for i := 0; i < jumlahElemen; i++ {
        fmt.Scan(&data[i])
    }

    // a. Menampilkan seluruh isi dari array
    fmt.Println("Isi lengkap array:", data)

    // b. Menampilkan elemen pada indeks ganjil
    fmt.Print("Elemen di indeks ganjil: ")
    for i := 1; i < len(data); i += 2 {
        fmt.Print(data[i], " ")
    }
    fmt.Println()

    // c. Menampilkan elemen pada indeks genap
    fmt.Print("Elemen di indeks genap: ")
    for i := 0; i < len(data); i += 2 {
        fmt.Print(data[i], " ")
    }
    fmt.Println()

    // d. Menampilkan elemen pada indeks kelipatan bilangan tertentu
    var kelipatan int
    fmt.Print("Masukkan nilai kelipatan untuk indeks: ")
    fmt.Scan(&kelipatan)

    fmt.Printf("Elemen pada indeks kelipatan %d: ", kelipatan)
    for i := kelipatan; i < len(data); i += kelipatan {
        fmt.Print(data[i], " ")
    }
    fmt.Println()

    // e. Menghapus elemen pada indeks tertentu
    var posisi int
    fmt.Print("Masukkan indeks yang ingin dihapus: ")
    fmt.Scan(&posisi)
    if posisi >= 0 && posisi < len(data) {
        data = hapusItem(data, posisi)
        fmt.Println("Isi array setelah penghapusan:", data)
    } else {
        fmt.Println("Indeks tidak valid!")
    }

    // f. Menampilkan rata-rata dari elemen-elemen dalam array
    fmt.Printf("Rata-rata elemen dalam array: %.2f\n", hitungRatarata_2311102181(data))

    // g. Menampilkan standar deviasi dari elemen-elemen dalam array
    fmt.Printf("Standar deviasi elemen dalam array: %.2f\n", hitungStandarDeviasi(data))

    // h. Menghitung frekuensi dari elemen tertentu dalam array
    var elemenCari int
    fmt.Print("Masukkan elemen yang ingin dihitung frekuensinya: ")
    fmt.Scan(&elemenCari)
    fmt.Printf("Frekuensi kemunculan elemen %d dalam array: %d\n", elemenCari, hitungFrekuensi(data, elemenCari))
}


