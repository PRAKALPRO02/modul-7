package main

import (
	"fmt"
	"math"
)

func main() {
	var N int
	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scan(&N)

	// Membuat array dengan kapasitas N
	arr := make([]int, N)

	// Mengisi array dengan nilai
	fmt.Println("Masukkan elemen array:")
	for i := 0; i < N; i++ {
		fmt.Printf("Elemen [%d]: ", i)
		fmt.Scan(&arr[i])
	}

	// Menampilkan keseluruhan isi dari array
	fmt.Println("\nIsi keseluruhan array:")
	displayArray(arr)

	fmt.Println("\nElemen dengan indeks ganjil:")
	displayOddIndex(arr)

	fmt.Println("\nElemen dengan indeks genap:")
	displayEvenIndex(arr)

	var x int
	fmt.Print("\nMasukkan bilangan x untuk indeks kelipatan: ")
	fmt.Scan(&x)
	fmt.Printf("\nElemen dengan indeks kelipatan %d:\n", x)
	displayMultipleOfX(arr, x)

	var indexToRemove int
	fmt.Print("\nMasukkan indeks yang ingin dihapus: ")
	fmt.Scan(&indexToRemove)
	arr = removeElement(arr, indexToRemove)
	fmt.Println("\nIsi array setelah penghapusan:")
	displayArray(arr)

	average := calculateAverage(arr)
	fmt.Printf("\nRata-rata: %.2f\n", average)

	stdDev := calculateStdDev(arr, average)
	fmt.Printf("Standar deviasi: %.2f\n", stdDev)

	var numberToFind int
	fmt.Print("\nMasukkan bilangan untuk mencari frekuensinya: ")
	fmt.Scan(&numberToFind)
	frequency := countFrequency(arr, numberToFind)
	fmt.Printf("Frekuensi dari bilangan %d: %d\n", numberToFind, frequency)
}

func displayArray(arr []int) {
	for _, v := range arr {
		fmt.Println(v)
	}
}

func displayOddIndex(arr []int) {
	for i := 1; i < len(arr); i += 2 {
		fmt.Println(arr[i])
	}
}

func displayEvenIndex(arr []int) {
	for i := 0; i < len(arr); i += 2 {
		fmt.Println(arr[i])
	}
}

func displayMultipleOfX(arr []int, x int) {
	for i := 0; i < len(arr); i++ {
		if i%x == 0 {
			fmt.Println(arr[i])
		}
	}
}

func removeElement(arr []int, index int) []int {
	return append(arr[:index], arr[index+1:]...)
}

func calculateAverage(arr []int) float64 {
	if len(arr) == 0 {
		return 0
	}
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return float64(sum) / float64(len(arr))
}

func calculateStdDev(arr []int, average float64) float64 {
	if len(arr) == 0 {
		return 0
	}
	var sum float64
	for _, v := range arr {
		sum += math.Pow(float64(v)-average, 2)
	}
	return math.Sqrt(sum / float64(len(arr)))
}

func countFrequency(arr []int, number int) int {
	count := 0
	for _, v := range arr {
		if v == number {
			count++
		}
	}
	return count
}
