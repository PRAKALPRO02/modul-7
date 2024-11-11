package main

import (
    "fmt"
    "math"
    "sort"
)

func tampilkanArray(arr []int) {
    fmt.Println("Isi array:", arr)
}

func tampilkanIndeksGanjil(arr []int) {
    fmt.Print("Elemen dengan indeks ganjil: ")
    for i := 0; i < len(arr); i += 2 { 
        fmt.Print(arr[i], " ")
    }
    fmt.Println()
}

func tampilkanIndeksGenap(arr []int) {
    fmt.Print("Elemen dengan indeks genap: ")
    for i := 1; i < len(arr); i += 2 { 
        fmt.Print(arr[i], " ")
    }
    fmt.Println()
}

func tampilkanKelipatanIndeks(arr []int, x int) {
    fmt.Printf("Elemen dengan indeks kelipatan %d: ", x)
    for i := x; i < len(arr); i += x {
        fmt.Print(arr[i], " ")
    }
    fmt.Println()
}

func hapusElemen(arr []int, index int) []int {
    if index < 0 || index >= len(arr) {
        fmt.Println("Indeks tidak valid!")
        return arr
    }
    return append(arr[:index], arr[index+1:]...)
}

func hitungRataRata(arr []int) float64 {
    total := 0
    for _, nilai := range arr {
        total += nilai
    }
    return float64(total) / float64(len(arr))
}

func hitungStandarDeviasi(arr []int) float64 {
    rata := hitungRataRata(arr)
    var total float64
    for _, nilai := range arr {
        total += math.Pow(float64(nilai)-rata, 2)
    }
    return math.Sqrt(total / float64(len(arr)))
}

func tampilkanFrekuensi(arr []int) {
    frekuensi := make(map[int]int)
    for _, nilai := range arr {
        frekuensi[nilai]++
    }
    
    keys := make([]int, 0, len(frekuensi))
    for k := range frekuensi {
        keys = append(keys, k)
    }
    sort.Ints(keys)

    fmt.Println("Frekuensi elemen dalam array:")
    for _, k := range keys {
        fmt.Printf("%d: %d kali\n", k, frekuensi[k])
    }
}

func main() {
    var n, x, index int
    fmt.Print("Masukkan jumlah elemen: ")
    fmt.Scan(&n)

    arr := make([]int, n)
    fmt.Println("Masukkan elemen-elemen array:")
    for i := 0; i < n; i++ {
        fmt.Scan(&arr[i])
    }

    tampilkanArray(arr)
    tampilkanIndeksGanjil(arr)
    tampilkanIndeksGenap(arr)

    fmt.Print("Masukkan nilai x: ")
    fmt.Scan(&x)
    tampilkanKelipatanIndeks(arr, x)

    fmt.Print("Masukkan indeks yang ingin dihapus: ")
    fmt.Scan(&index)
    arr = hapusElemen(arr, index)
    tampilkanArray(arr)

    rataRata := hitungRataRata(arr)
    fmt.Printf("Rata-rata: %.2f\n", rataRata)

    standarDeviasi := hitungStandarDeviasi(arr)
    fmt.Printf("Standar deviasi: %.2f\n", standarDeviasi)

    tampilkanFrekuensi(arr)
}