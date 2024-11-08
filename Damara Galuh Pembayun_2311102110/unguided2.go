package main

import (
        "fmt"
        "math"
)

func main() {
        var N int
        fmt.Print("Masukkan jumlah elemen array: ")
        fmt.Scan(&N)

        // Membuat array dengan ukuran N
        arr := make([]int, N)

        // Meminta input nilai untuk setiap elemen array
        fmt.Println("Masukkan nilai-nilai array:")
        for i := 0; i < N; i++ {
                fmt.Scan(&arr[i])
        }

        // Menampilkan seluruh isi array
        fmt.Println("Seluruh isi array:")
        for _, v := range arr {
                fmt.Print(v, " ")
        }
        fmt.Println()

        // Menampilkan elemen dengan indeks ganjil
        fmt.Println("Elemen dengan indeks ganjil:")
        for i := 1; i < N; i += 2 {
                fmt.Print(arr[i], " ")
        }
        fmt.Println()

        // Menampilkan elemen dengan indeks genap
        fmt.Println("Elemen dengan indeks genap:")
        for i := 0; i < N; i += 2 {
                fmt.Print(arr[i], " ")
        }
        fmt.Println()

        // Menampilkan elemen dengan indeks kelipatan x
        var x int
        fmt.Print("Masukkan nilai x: ")
        fmt.Scan(&x)
        fmt.Println("Elemen dengan indeks kelipatan", x)
        for i := 0; i < N; i++ {
                if i%x == 0 {
                        fmt.Print(arr[i], " ")
                }
        }
        fmt.Println()

        // Menghapus elemen pada indeks tertentu
        var index int
        fmt.Print("Masukkan indeks yang ingin dihapus: ")
        fmt.Scan(&index)
        arr = append(arr[:index], arr[index+1:]...)
        fmt.Println("Array setelah penghapusan:")
        for _, v := range arr {
                fmt.Print(v, " ")
        }
        fmt.Println()

        // Menghitung rata-rata
        var sum int
        for _, v := range arr {
                sum += v
        }
        rataRata := float64(sum) / float64(len(arr))
        fmt.Println("Rata-rata:", rataRata)

        // Menghitung standar deviasi
        var variance float64
        for _, v := range arr {
                variance += math.Pow(float64(v)-rataRata, 2)
        }
        variance /= float64(len(arr))
        standarDeviasi := math.Sqrt(variance)
        fmt.Println("Standar deviasi:", standarDeviasi)

        // Menghitung frekuensi suatu bilangan
        var bilangan int
        fmt.Print("Masukkan bilangan yang ingin dicari frekuensinya: ")
        fmt.Scan(&bilangan)
        var frekuensi int
        for _, v := range arr {
                if v == bilangan {
                        frekuensi++
                }
        }
        fmt.Println("Frekuensi", bilangan, "adalah", frekuensi)
}