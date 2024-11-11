package main

import (
	"fmt"
	"math"
)

func displayarr(arr []int) {
	fmt.Println("element array:", arr)
}

func bilangan_ganjil(arr []int) {
	fmt.Print("Elemen ganjil pada array: ")
	for i := 1; i < len(arr); i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}

func display_elementindx(arr []int) {
	fmt.Print("Element genap pada array: ")
	for i := 0; i < len(arr); i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}

func display_kelipatan(arr []int, x int) {
	fmt.Printf("Kelipatan dari element index %d: ", x)
	for i := x; i < len(arr); i += x {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}

func hapus_element(arr []int, index int) []int {
	if index < 0 || index >= len(arr) {
		fmt.Println("Index tidk ditemukan")
		return arr
	}
	fmt.Printf(" index %d telah dihapus\n", index)
	return append(arr[:index], arr[index+1:]...)
}

func nilai_ratarata(arr []int) float64 {
	sum := 0
	for _, value := range arr {
		sum += value
	}
	return float64(sum) / float64(len(arr))
}

func rata_bilangan(arr []int) float64 {
	mean := nilai_ratarata(arr)
	sum := 0.0
	for _, value := range arr {
		sum += math.Pow(float64(value)-mean, 2)
	}
	return math.Sqrt(sum / float64(len(arr)))
}

func simpangbaku(arr []int, number int) int {
	frequency := 0
	for _, value := range arr {
		if value == number {
			frequency++
		}
	}
	return frequency
}

func main() {
	var n, x, index, target int

	fmt.Print("Masukan angka untuk element di dalam array: ")
	fmt.Scan(&n)

	arr := make([]int, n)
	fmt.Println("masukan: ")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Println("\nTABEL ARRAY")
	fmt.Println("a. Tampilkan seluruh array")
	fmt.Println("b. Tampilkan array dengan ideks ganjil")
	fmt.Println("c. Tampilkan array dengan indeks genap")
	fmt.Println("d. Kelipatan")
	fmt.Println("e. Hapus element")
	fmt.Println("f. Menghitung rata-rata")
	fmt.Println("g. Mengitung standar deviasi atau simpang baku")
	fmt.Println("h. Menghitung fekuensi dari suatu bilangan")

	var pilih string
	fmt.Print("Masukan pilihan anda (a-h): ")
	fmt.Scan(&pilih)

	switch pilih {
	case "a":
		displayarr(arr)
	case "b":
		bilangan_ganjil(arr)
	case "c":
		display_elementindx(arr)
	case "d":
		fmt.Print("Masukan nilai kelipatan x: ")
		fmt.Scan(&x)
		display_kelipatan(arr, x)
	case "e":
		fmt.Print("Index yang ingin di hapus: ")
		fmt.Scan(&index)
		arr = hapus_element(arr, index)
		displayarr(arr)
	case "f":
		average := nilai_ratarata(arr)
		fmt.Printf("Rata-rata dari element: %.2f\n", average)
	case "g":
		stdDev := rata_bilangan(arr)
		fmt.Printf("standar atau simpang bakunya: %.2f\n", stdDev)
	case "h":
		fmt.Print("Masukan angka untuk mencari freuensi: ")
		fmt.Scan(&target)
		frequency := simpangbaku(arr, target)
		fmt.Printf("Frekuensi dari %d: %d\n", target, frequency)
	default:
		fmt.Println("ERRORR")
	}
}
