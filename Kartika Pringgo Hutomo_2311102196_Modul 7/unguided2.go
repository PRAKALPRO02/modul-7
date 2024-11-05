package main

import (
	"fmt"
	"math"
)

func main() {
	var N int
	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scan(&N)

	arr := make([]int, N)

	fmt.Println("Masukkan elemen array:")
	for i := 0; i < N; i++ {
		fmt.Printf("Elemen [%d]: ", i)
		fmt.Scan(&arr[i])
	}
	
	fmt.Println("\nIsi seluruh array:")
	tampilkanArray(arr)

	fmt.Println("\nElemen pada indeks ganjil:")
	tampilkanIndeksGanjil(arr)

	fmt.Println("\nElemen pada indeks genap:")
	tampilkanIndeksGenap(arr)

	var x int
	fmt.Print("\nMasukkan angka x untuk indeks kelipatan: ")
	fmt.Scan(&x)
	fmt.Printf("\nElemen pada indeks kelipatan %d:\n", x)
	tampilkanKelipatanX(arr, x)

	var indeksHapus int
	fmt.Print("\nMasukkan indeks yang ingin dihapus: ")
	fmt.Scan(&indeksHapus)
	arr = hapusElemen(arr, indeksHapus)
	fmt.Println("\nIsi array setelah penghapusan:")
	tampilkanArray(arr)

	rataRata := hitungRataRata(arr)
	fmt.Printf("\nRata-rata: %.2f\n", rataRata)

	stdDev := hitungStandarDeviasi(arr, rataRata)
	fmt.Printf("Standar deviasi: %.2f\n", stdDev)

	var angkaCari int
	fmt.Print("\nMasukkan angka untuk mencari frekuensinya: ")
	fmt.Scan(&angkaCari)
	frekuensi := hitungFrekuensi(arr, angkaCari)
	fmt.Printf("Frekuensi dari angka %d: %d\n", angkaCari, frekuensi)
}

func tampilkanArray(arr []int) {
	for _, v := range arr {
		fmt.Println(v)
	}
}

func tampilkanIndeksGanjil(arr []int) {
	for i := 0; i < len(arr); i += 2 {
		fmt.Println(arr[i])
	}
}

func tampilkanIndeksGenap(arr []int) {
	for i := 1; i < len(arr); i += 2 {
		fmt.Println(arr[i])
	}
}

func tampilkanKelipatanX(arr []int, x int) {
	for i := 0; i < len(arr); i++ {
		if i%x == 0 {
			fmt.Println(arr[i])
		}
	}
}

func hapusElemen(arr []int, index int) []int {
	return append(arr[:index], arr[index+1:]...)
}

func hitungRataRata(arr []int) float64 {
	if len(arr) == 0 {
		return 0
	}
	total := 0
	for _, v := range arr {
		total += v
	}
	return float64(total) / float64(len(arr))
}

func hitungStandarDeviasi(arr []int, rataRata float64) float64 {
	if len(arr) == 0 {
		return 0
	}
	var total float64
	for _, v := range arr {
		total += math.Pow(float64(v)-rataRata, 2)
	}
	return math.Sqrt(total / float64(len(arr)))
}

func hitungFrekuensi(arr []int, angka int) int {
	jumlah := 0
	for _, v := range arr {
		if v == angka {
			jumlah++
		}
	}
	return jumlah
}
