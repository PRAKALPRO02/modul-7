package main

import (
	"fmt"
	"math"
)

// Fungsi untuk menghitung rata-rata
func rataRata(arr []int) float64 {
	total := 0
	for _, v := range arr {
		total += v
	}
	return float64(total) / float64(len(arr))
}

// Fungsi untuk menghitung standar deviasi
func standarDeviasi(arr []int, mean float64) float64 {
	var total float64
	for _, v := range arr {
		total += math.Pow(float64(v)-mean, 2)
	}
	return math.Sqrt(total / float64(len(arr)))
}

// Fungsi untuk menghitung frekuensi suatu bilangan
func frekuensi(arr []int, bilangan int) int {
	count := 0
	for _, v := range arr {
		if v == bilangan {
			count++
		}
	}
	return count
}

func main() {
	var N int
	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scan(&N)

	arr := make([]int, N)

	// Input elemen array
	fmt.Println("Masukkan elemen array:")
	for i := 0; i < N; i++ {
		fmt.Printf("Elemen ke-%d: ", i+1)
		fmt.Scan(&arr[i])
	}

	// Menampilkan seluruh isi array
	fmt.Println("Isi array:", arr)

	// Menampilkan elemen dengan indeks ganjil
	fmt.Print("Elemen dengan indeks ganjil: ")
	for i := 1; i < N; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	// Menampilkan elemen dengan indeks genap
	fmt.Print("Elemen dengan indeks genap: ")
	for i := 0; i < N; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	// Input bilangan x untuk kelipatan
	var x int
	fmt.Print("Masukkan bilangan x untuk kelipatan: ")
	fmt.Scan(&x)

	// Menampilkan elemen dengan indeks kelipatan x
	fmt.Print("Elemen dengan indeks kelipatan ", x, ": ")
	for i := 0; i < N; i++ {
		if i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	// Menghapus elemen pada indeks tertentu
	var indeksHapus int
	fmt.Print("Masukkan indeks yang ingin dihapus: ")
	fmt.Scan(&indeksHapus)

	if indeksHapus >= 0 && indeksHapus < N {
		arr = append(arr[:indeksHapus], arr[indeksHapus+1:]...)
		fmt.Println("Array setelah penghapusan:", arr)
	} else {
		fmt.Println("Indeks tidak valid.")
	}

	// Menghitung dan menampilkan rata-rata
	mean := rataRata(arr)
	fmt.Printf("Rata-rata: %.2f\n", mean)

	// Menghitung dan menampilkan standar deviasi
	stdDev := standarDeviasi(arr, mean)
	fmt.Printf("Standar deviasi: %.2f\n", stdDev)

	// Input bilangan untuk frekuensi
	var bilangan int
	fmt.Print("Masukkan bilangan untuk menghitung frekuensi: ")
	fmt.Scan(&bilangan)
	freq := frekuensi(arr, bilangan)
	fmt.Printf("Frekuensi bilangan %d: %d\n", bilangan, freq)
}
