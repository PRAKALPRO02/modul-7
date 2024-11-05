// Ben Waiz Pintus W
// 2311102169
package main

import (
	"fmt"
	"math"
)

// Fungsi untuk menampilkan seluruh elemen array
func tampilkanSeluruhArray(arr []int) {
	fmt.Println("Seluruh elemen array:", arr)
}

// Fungsi untuk menampilkan elemen array dengan indeks ganjil
func tampilkanIndeksGanjil(arr []int) {
	fmt.Print("Elemen dengan indeks ganjil: ")
	for i := 0; i < len(arr); i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}

// Fungsi untuk menampilkan elemen array dengan indeks genap
func tampilkanIndeksGenap(arr []int) {
	fmt.Print("Elemen dengan indeks genap: ")
	for i := 1; i < len(arr); i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}

// Fungsi untuk menampilkan elemen array dengan indeks kelipatan x
func tampilkanIndeksKelipatan(arr []int, x int) {
	fmt.Printf("Elemen dengan indeks kelipatan %d: ", x)
	for i := 0; i < len(arr); i++ {
		if i%x == 0 { // Periksa apakah indeks i adalah kelipatan dari x
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()
}

// Fungsi untuk menghapus elemen array pada indeks tertentu
func hapusElemen(arr []int, indeks int) []int {
	return append(arr[:indeks], arr[indeks+1:]...)
}

// Fungsi untuk menghitung rata-rata array
func rataRata(arr []int) float64 {
	total := 0
	for _, v := range arr {
		total += v
	}
	return float64(total) / float64(len(arr))
}

// Fungsi untuk menghitung standar deviasi array
func standarDeviasi(arr []int) float64 {
	rata := rataRata(arr)
	var totalKuadrat float64
	for _, v := range arr {
		totalKuadrat += math.Pow(float64(v)-rata, 2)
	}
	return math.Sqrt(totalKuadrat / float64(len(arr)))
}

// Fungsi untuk menghitung frekuensi elemen tertentu dalam array
func frekuensi(arr []int, elemen int) int {
	count := 0
	for _, v := range arr {
		if v == elemen {
			count++
		}
	}
	return count
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

	tampilkanSeluruhArray(arr)
	tampilkanIndeksGanjil(arr)
	tampilkanIndeksGenap(arr)

	var x int
	fmt.Print("Masukkan bilangan x untuk indeks kelipatan: ")
	fmt.Scan(&x)
	tampilkanIndeksKelipatan(arr, x)

	var indeksHapus int
	fmt.Print("Masukkan indeks yang akan dihapus: ")
	fmt.Scan(&indeksHapus)
	arr = hapusElemen(arr, indeksHapus)
	fmt.Println("Array setelah penghapusan elemen:")
	tampilkanSeluruhArray(arr)

	fmt.Printf("Rata-rata array: %.2f\n", rataRata(arr))
	fmt.Printf("Standar deviasi array: %.2f\n", standarDeviasi(arr))

	var elemenCari int
	fmt.Print("Masukkan elemen untuk mencari frekuensi: ")
	fmt.Scan(&elemenCari)
	fmt.Printf("Frekuensi elemen %d: %d kali\n", elemenCari, frekuensi(arr, elemenCari))
}
