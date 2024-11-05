package main

// NAMA : CHRIST DANIEL SANTOSO
// NIM : 2311102305

import (
	"fmt"
	"math"
)

func tampilkanArray(arr []int) {
	fmt.Println("Seluruh elemen array:", arr)
}

func tampilkanElemenIndeksGanjil(arr []int) {
	fmt.Print("Elemen dengan indeks ganjil: ")
	for i := 0; i < len(arr); i++ {
		if i%2 == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()
}

func tampilkanElemenIndeksGenap(arr []int) {
	fmt.Print("Elemen dengan indeks genap: ")
	for i := 1; i < len(arr); i++ {
		if i%2 != 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()
}

func tampilkanElemenIndeksKelipatan(arr []int, x int) {
	fmt.Printf("Elemen dengan indeks kelipatan %d: ", x)
	for i := 0; i < len(arr); i++ {
		if i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()
}

func hapusPadaIndeks(arr []int, idx int) []int {
	return append(arr[:idx], arr[idx+1:]...)
}

func hitungRataRata(arr []int) float64 {
	total := 0
	for _, v := range arr {
		total += v
	}
	return float64(total) / float64(len(arr))
}

func hitungStandarDeviasi(arr []int) float64 {
	rata := hitungRataRata(arr)
	var totalKuadrat float64
	for _, v := range arr {
		totalKuadrat += math.Pow(float64(v)-rata, 2)
	}
	return math.Sqrt(totalKuadrat / float64(len(arr)))
}

func hitungFrekuensi(arr []int, elemen int) int {
	jumlah := 0
	for _, v := range arr {
		if v == elemen {
			jumlah++
		}
	}
	return jumlah
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scan(&n)

	arr := make([]int, n)
	fmt.Println("Masukkan elemen-elemen array:")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	tampilkanArray(arr)
	tampilkanElemenIndeksGanjil(arr)
	tampilkanElemenIndeksGenap(arr)

	var x int
	fmt.Print("Masukkan bilangan x untuk indeks kelipatan: ")
	fmt.Scan(&x)
	tampilkanElemenIndeksKelipatan(arr, x)

	var idxHapus int
	fmt.Print("Masukkan indeks yang akan dihapus: ")
	fmt.Scan(&idxHapus)
	arr = hapusPadaIndeks(arr, idxHapus)
	fmt.Println("Array setelah penghapusan elemen:")
	tampilkanArray(arr)

	fmt.Printf("Rata-rata array: %.2f\n", hitungRataRata(arr))
	fmt.Printf("Standar deviasi array: %.2f\n", hitungStandarDeviasi(arr))

	var elemenCari int
	fmt.Print("Masukkan elemen untuk mencari frekuensi: ")
	fmt.Scan(&elemenCari)
	fmt.Printf("Frekuensi elemen %d: %d kali\n", elemenCari, hitungFrekuensi(arr, elemenCari))
}
