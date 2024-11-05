package main

import (
	"fmt"
	"math"
)

// Fungsi untuk menghapus elemen pada indeks tertentu
func hapusElemen(arr []int, index int) []int {
	return append(arr[:index], arr[index+1:]...)
}

// Fungsi untuk menghitung rata-rata
func rataRata(arr []int) float64 {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return float64(sum) / float64(len(arr))
}

// Fungsi untuk menghitung standar deviasi
func standarDeviasi(arr []int) float64 {
	mean := rataRata(arr)
	var variance float64
	for _, val := range arr {
		variance += math.Pow(float64(val)-mean, 2)
	}
	return math.Sqrt(variance / float64(len(arr)))
}

// Fungsi untuk menghitung frekuensi elemen tertentu dalam array
func frekuensi(arr []int, x int) int {
	count := 0
	for _, val := range arr {
		if val == x {
			count++
		}
	}
	return count
}

func main() {
	// Memasukkan jumlah elemen dan isi dari array
	var n int
	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scan(&n)

	arr := make([]int, n)
	fmt.Println("Masukkan elemen-elemen array:")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	// a. Menampilkan keseluruhan isi dari array
	fmt.Println("Keseluruhan isi array:", arr)

	// b. Menampilkan elemen-elemen array dengan indeks ganjil
	fmt.Print("Elemen dengan indeks ganjil: ")
	for i := 1; i < len(arr); i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	// c. Menampilkan elemen-elemen array dengan indeks genap
	fmt.Print("Elemen dengan indeks genap: ")
	for i := 0; i < len(arr); i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	// d. Menampilkan elemen array dengan indeks kelipatan bilangan x
	var x int
	fmt.Print("Masukkan nilai x untuk indeks kelipatan: ")
	fmt.Scan(&x)

	fmt.Print("Elemen pada indeks kelipatan ", x, ": ")
	for i := x; i < len(arr); i += x {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	// e. Menghapus elemen array pada indeks tertentu
	var index int
	fmt.Print("Masukkan indeks yang ingin dihapus: ")
	fmt.Scan(&index)
	if index >= 0 && index < len(arr) {
		arr = hapusElemen(arr, index)
		fmt.Println("Isi array setelah penghapusan:", arr)
	} else {
		fmt.Println("Indeks tidak valid!")
	}

	// f. Menampilkan rata-rata dari bilangan yang ada di dalam array
	fmt.Printf("Rata-rata bilangan di array: %.2f\n", rataRata(arr))

	// g. Menampilkan standar deviasi dari bilangan yang ada di dalam array
	fmt.Printf("Standar deviasi bilangan di array: %.2f\n", standarDeviasi(arr))

	// h. Menampilkan frekuensi dari suatu bilangan tertentu di dalam array
	var cari int
	fmt.Print("Masukkan bilangan yang ingin dihitung frekuensinya: ")
	fmt.Scan(&cari)
	fmt.Printf("Frekuensi bilangan %d dalam array: %d\n", cari, frekuensi(arr, cari))
}
