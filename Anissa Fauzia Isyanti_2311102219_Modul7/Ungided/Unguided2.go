package main

import (
	"fmt"
	"math"
)

func main() {
	var N int
	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scan(&N)

	// Inisialisasi array
	array := make([]int, N)

	// Mengisi array dengan nilai dari pengguna
	for i := 0; i < N; i++ {
		fmt.Printf("Masukkan elemen ke-%d: ", i)
		fmt.Scan(&array[i])
	}

	// Menampilkan keseluruhan isi dari array
	fmt.Println("\nIsi array:")
	fmt.Println(array)

	// Menampilkan elemen dengan indeks ganjil
	fmt.Println("\nElemen dengan indeks ganjil:")
	for i := 1; i < N; i += 2 {
		fmt.Print(array[i], " ")
	}

	// Menampilkan elemen dengan indeks genap
	fmt.Println("\n\nElemen dengan indeks genap:")
	for i := 0; i < N; i += 2 {
		fmt.Print(array[i], " ")
	}

	// Menampilkan elemen dengan indeks kelipatan bilangan x
	var x int
	fmt.Print("\n\nMasukkan bilangan x: ")
	fmt.Scan(&x)
	fmt.Println("Elemen dengan indeks kelipatan", x, ":")
	for i := 0; i < N; i++ {
		if i%x == 0 {
			fmt.Print(array[i], " ")
		}
	}

	// Menghapus elemen pada indeks tertentu
	var indeksHapus int
	fmt.Print("\n\nMasukkan indeks yang akan dihapus: ")
	fmt.Scan(&indeksHapus)
	array = append(array[:indeksHapus], array[indeksHapus+1:]...)
	fmt.Println("Array setelah dihapus elemen pada indeks", indeksHapus, ":")
	fmt.Println(array)

	// Menghitung dan menampilkan rata-rata
	var total int
	for _, value := range array {
		total += value
	}
	rataRata := float64(total) / float64(len(array))
	fmt.Printf("\nRata-rata: %.2f\n", rataRata)

	// Menghitung dan menampilkan standar deviasi
	var variance float64
	for _, value := range array {
		variance += math.Pow(float64(value)-rataRata, 2)
	}
	variance /= float64(len(array))
	standarDeviasi := math.Sqrt(variance)
	fmt.Printf("Standar deviasi: %.2f\n", standarDeviasi)

	// Menampilkan frekuensi dari suatu bilangan tertentu
	var bilangan int
	fmt.Print("\nMasukkan bilangan yang frekuensinya ingin dihitung: ")
	fmt.Scan(&bilangan)
	frekuensi := 0
	for _, value := range array {
		if value == bilangan {
			frekuensi++
		}
	}
	fmt.Printf("Frekuensi bilangan %d: %d\n", bilangan, frekuensi)
}
