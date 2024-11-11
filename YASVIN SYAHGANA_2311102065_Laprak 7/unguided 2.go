package main

import (
	"fmt"
	"math"
)

func main() {
	var n, x, delIndex, searchValue int

	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scan(&n)
	array := make([]int, n)

	fmt.Println("Masukkan elemen-elemen array:")
	for i := 0; i < n; i++ {
		fmt.Printf("Elemen ke-%d: ", i+1)
		fmt.Scan(&array[i])
	}

	fmt.Println("a. Isi keseluruhan array:", array)

	fmt.Print("b. Elemen dengan indeks ganjil: ")
	for i := 1; i < n; i += 2 {
		fmt.Print(array[i], " ")
	}
	fmt.Println()

	fmt.Print("c. Elemen dengan indeks genap: ")
	for i := 0; i < n; i += 2 {
		fmt.Print(array[i], " ")
	}
	fmt.Println()
	fmt.Print("Masukkan nilai x untuk menampilkan elemen dengan indeks kelipatan x: ")
	fmt.Scan(&x)
	fmt.Printf("d. Elemen dengan indeks kelipatan %d: ", x)
	for i := x; i < n; i += x {
		fmt.Print(array[i], " ")
	}
	fmt.Println()
	fmt.Print("Masukkan indeks elemen yang ingin dihapus: ")
	fmt.Scan(&delIndex)
	if delIndex >= 0 && delIndex < n {
		array = append(array[:delIndex], array[delIndex+1:]...)
		fmt.Println("e. Isi array setelah penghapusan:", array)
	} else {
		fmt.Println("Indeks tidak valid.")
	}

	sum := 0
	for _, num := range array {
		sum += num
	}
	average := float64(sum) / float64(len(array))
	fmt.Printf("f. Rata-rata elemen: %.2f\n", average)

	var deviationSum float64
	for _, num := range array {
		deviationSum += math.Pow(float64(num)-average, 2)
	}
	standardDeviation := math.Sqrt(deviationSum / float64(len(array)))
	fmt.Printf("g. Standar deviasi: %.2f\n", standardDeviation)

	fmt.Print("Masukkan nilai untuk menghitung frekuensinya di dalam array: ")
	fmt.Scan(&searchValue)
	frequency := 0
	for _, num := range array {
		if num == searchValue {
			frequency++
		}
	}
	fmt.Printf("h. Frekuensi bilangan %d di dalam array: %d\n", searchValue, frequency)
}
