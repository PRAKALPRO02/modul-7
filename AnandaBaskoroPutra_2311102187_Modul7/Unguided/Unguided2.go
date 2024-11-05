package main

import (
	"fmt"
	"math"
)

// Fungsi untuk menampilkan seluruh isi array
func tampilkanSeluruhArray(arr []int) {
	fmt.Println("Isi array:", arr)
}

// Fungsi untuk menampilkan elemen-elemen array dengan indeks ganjil
func tampilkanIndeksGanjil(arr []int) {
	fmt.Print("Elemen dengan indeks ganjil: ")
	for i := 1; i < len(arr); i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}

// Fungsi untuk menampilkan elemen-elemen array dengan indeks genap
func tampilkanIndeksGenap(arr []int) {
	fmt.Print("Elemen dengan indeks genap: ")
	for i := 0; i < len(arr); i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}

// Fungsi untuk menampilkan elemen-elemen array dengan indeks kelipatan x
func tampilkanIndeksKelipatan(arr []int, x int) {
	fmt.Printf("Elemen dengan indeks kelipatan %d: ", x)
	for i := x; i < len(arr); i += x {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}

// Fungsi untuk menghapus elemen pada indeks tertentu
func hapusElemen(arr []int, index int) []int {
	if index < 0 || index >= len(arr) {
		fmt.Println("Indeks tidak valid.")
		return arr
	}
	return append(arr[:index], arr[index+1:]...)
}

// Fungsi untuk menghitung rata-rata elemen dalam array
func hitungRataRata(arr []int) float64 {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return float64(sum) / float64(len(arr))
}

// Fungsi untuk menghitung standar deviasi elemen dalam array
func hitungStandarDeviasi(arr []int) float64 {
	rataRata := hitungRataRata(arr)
	var sum float64
	for _, v := range arr {
		sum += math.Pow(float64(v)-rataRata, 2)
	}
	varian := sum / float64(len(arr))
	return math.Sqrt(varian)
}

// Fungsi untuk menghitung frekuensi dari setiap elemen dalam array
func hitungFrekuensi(arr []int) map[int]int {
	frekuensi := make(map[int]int)
	for _, v := range arr {
		frekuensi[v]++
	}
	return frekuensi
}

func main() {
	// Input jumlah elemen dan nilai-nilai elemen array
	var n int
	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scan(&n)

	arr := make([]int, n)
	fmt.Println("Masukkan elemen-elemen array:")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	// Menampilkan seluruh isi array
	tampilkanSeluruhArray(arr)

	// Menampilkan elemen dengan indeks ganjil
	tampilkanIndeksGanjil(arr)

	// Menampilkan elemen dengan indeks genap
	tampilkanIndeksGenap(arr)

	// Menampilkan elemen dengan indeks kelipatan x
	var x int
	fmt.Print("Masukkan nilai x untuk indeks kelipatan: ")
	fmt.Scan(&x)
	tampilkanIndeksKelipatan(arr, x)

	// Menghapus elemen pada indeks tertentu
	var index int
	fmt.Print("Masukkan indeks elemen yang ingin dihapus: ")
	fmt.Scan(&index)
	arr = hapusElemen(arr, index)
	fmt.Println("Array setelah elemen dihapus:")
	tampilkanSeluruhArray(arr)

	// Menampilkan rata-rata dari elemen-elemen dalam array
	rataRata := hitungRataRata(arr)
	fmt.Printf("Rata-rata elemen array: %.2f\n", rataRata)

	// Menampilkan standar deviasi dari elemen-elemen dalam array
	standarDeviasi := hitungStandarDeviasi(arr)
	fmt.Printf("Standar deviasi elemen array: %.2f\n", standarDeviasi)

	// Menampilkan frekuensi dari setiap elemen dalam array
	frekuensi := hitungFrekuensi(arr)
	fmt.Println("Frekuensi dari setiap elemen:")
	for key, value := range frekuensi {
		fmt.Printf("Elemen %d muncul %d kali\n", key, value)
	}
}
