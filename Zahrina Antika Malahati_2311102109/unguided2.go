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

	// Meminta pengguna untuk mengisi elemen array
	fmt.Println("Masukkan elemen array:")
	for i := 0; i < N; i++ {
		fmt.Scan(&arr[i])
	}

	// Menampilkan keseluruhan isi array
	fmt.Println("Seluruh isi array:")
	fmt.Println(arr)

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
	fmt.Printf("Elemen dengan indeks kelipatan %d:\n", x)
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
	fmt.Println(arr)

	// Menghitung rata-rata
	var sum int
	for _, num := range arr {
		sum += num
	}
	rataRata := float64(sum) / float64(len(arr))
	fmt.Printf("Rata-rata: %.2f\n", rataRata)

	// Menghitung standar deviasi
	var variance float64
	for _, num := range arr {
		variance += math.Pow(float64(num)-rataRata, 2)
	}
	variance /= float64(len(arr))
	standarDeviasi := math.Sqrt(variance)
	fmt.Printf("Standar deviasi: %.2f\n", standarDeviasi)

	// Menghitung frekuensi suatu bilangan
	var bilangan int
	fmt.Print("Masukkan bilangan yang ingin dicari frekuensinya: ")
	fmt.Scan(&bilangan)
	var frekuensi int
	for _, num := range arr {
		if num == bilangan {
			frekuensi++
		}
	}
	fmt.Printf("Frekuensi %d: %d\n", bilangan, frekuensi)
}

// Zahrina Antika Malahati_2311102109
