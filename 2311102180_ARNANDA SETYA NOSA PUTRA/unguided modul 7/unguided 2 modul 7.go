package main

import (
	"fmt"
	"math"
)

func main() {
	var N int
	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scanln(&N)

	arr := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Printf("Masukkan nilai untuk elemen ke-%d: ", i)
		fmt.Scanln(&arr[i])
	}

	fmt.Println("a. Isi keseluruhan array:", arr)

	fmt.Print("b. Elemen dengan indeks ganjil: ")
	for i := 1; i < N; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	fmt.Print("c. Elemen dengan indeks genap: ")
	for i := 0; i < N; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	var x int
	fmt.Print("Masukkan nilai x untuk menampilkan elemen dengan indeks kelipatan x: ")
	fmt.Scanln(&x)
	fmt.Print("d. Elemen dengan indeks kelipatan ", x, ": ")
	for i := x; i < N; i += x {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	var delIndex int
	fmt.Print("Masukkan indeks elemen yang ingin dihapus: ")
	fmt.Scanln(&delIndex)
	if delIndex >= 0 && delIndex < N {
		arr = append(arr[:delIndex], arr[delIndex+1:]...)
	}
	fmt.Println("e. Isi array setelah penghapusan:", arr)

	sum := 0
	for _, val := range arr {
		sum += val
	}
	average := float64(sum) / float64(len(arr))
	fmt.Printf("f. Rata-rata elemen: %.2f\n", average)

	var deviationSum float64
	for _, val := range arr {
		deviationSum += math.Pow(float64(val)-average, 2)
	}
	standardDeviation := math.Sqrt(deviationSum / float64(len(arr)))
	fmt.Printf("g. Standar deviasi: %.2f\n", standardDeviation)

	var searchValue int
	fmt.Print("Masukkan bilangan untuk menghitung frekuensinya di dalam array: ")
	fmt.Scanln(&searchValue)
	frequency := 0
	for _, val := range arr {
		if val == searchValue {
			frequency++
		}
	}
	fmt.Printf("h. Frekuensi bilangan %d di dalam array: %d\n", searchValue, frequency)
}
