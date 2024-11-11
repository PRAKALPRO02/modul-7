package main

import (
	"fmt"
	"math"
)

func displayArray(arr []int) {
	fmt.Println("Isi array:", arr)
}

func displayElementsByIndex(arr []int, start int, step int, description string) {
	fmt.Printf("Elemen dengan indeks %s: ", description)
	for i := start; i < len(arr); i += step {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}

func deleteElementAtIndex(arr *[]int, index int) {
	if index >= 0 && index < len(*arr) {
		*arr = append((*arr)[:index], (*arr)[index+1:]...)
		displayArray(*arr)
	} else {
		fmt.Println("Indeks tidak valid.")
	}
}

func calculateAverage(arr []int) float64 {
	if len(arr) == 0 {
		return 0
	}
	sum := 0
	for _, value := range arr {
		sum += value
	}
	return float64(sum) / float64(len(arr))
}

func calculateStandardDeviation(arr []int, average float64) float64 {
	if len(arr) == 0 {
		return 0
	}
	var varianceSum float64
	for _, value := range arr {
		varianceSum += math.Pow(float64(value)-average, 2)
	}
	return math.Sqrt(varianceSum / float64(len(arr)))
}

func countFrequency(arr []int, num int) int {
	count := 0
	for _, value := range arr {
		if value == num {
			count++
		}
	}
	return count
}

func main() {
	var numElements int
	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scan(&numElements)

	arr := make([]int, numElements)
	for i := 0; i < numElements; i++ {
		fmt.Printf("Masukkan elemen ke-%d: ", i)
		fmt.Scan(&arr[i])
	}

	displayArray(arr)

	displayElementsByIndex(arr, 1, 2, "ganjil")
	displayElementsByIndex(arr, 0, 2, "genap")

	var x int
	fmt.Print("Masukkan bilangan x untuk kelipatan indeks: ")
	fmt.Scan(&x)
	displayElementsByIndex(arr, x, x, fmt.Sprintf("kelipatan %d", x))

	var index int
	fmt.Print("Masukkan indeks untuk menghapus elemen: ")
	fmt.Scan(&index)
	deleteElementAtIndex(&arr, index)

	average := calculateAverage(arr)
	fmt.Println("Rata-rata array:", average)
	stdDev := calculateStandardDeviation(arr, average)
	fmt.Println("Simpangan baku (standar deviasi):", stdDev)

	var num int
	fmt.Print("Masukkan bilangan untuk menghitung frekuensi: ")
	fmt.Scan(&num)
	fmt.Printf("Frekuensi bilangan %d: %d\n", num, countFrequency(arr, num))
}
