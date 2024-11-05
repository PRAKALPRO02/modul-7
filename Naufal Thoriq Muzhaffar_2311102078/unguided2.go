package main

import (
	"fmt"
	"math"
)

func main() {
	var n, x, index, value int
	fmt.Print("Masukkan jumlah elemen dalam array: ")
	fmt.Scan(&n)

	array := make([]int, n)
	fmt.Println("Masukkan elemen-elemen array:")
	for i := 0; i < n; i++ {
		fmt.Scan(&array[i])
	}

	fmt.Println("a. Seluruh isi array:", array)
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

	fmt.Print("Masukkan nilai x untuk kelipatan indeks: ")
	fmt.Scan(&x)
	fmt.Print("d. Elemen dengan indeks kelipatan ", x, ": ")
	for i := 0; i < n; i += x {
		fmt.Print(array[i], " ")
	}
	fmt.Println()

	fmt.Print("Masukkan indeks untuk menghapus elemen: ")
	fmt.Scan(&index)
	if index >= 0 && index < n {
		array = append(array[:index], array[index+1:]...)
		fmt.Println("e. Array setelah penghapusan:", array)
	} else {
		fmt.Println("Indeks tidak valid.")
	}

	sum := 0
	for _, num := range array {
		sum += num
	}
	avg := float64(sum) / float64(len(array))
	fmt.Printf("f. Rata-rata: %.2f\n", avg)

	varianceSum := 0.0
	for _, num := range array {
		varianceSum += math.Pow(float64(num)-avg, 2)
	}
	stdDev := math.Sqrt(varianceSum / float64(len(array)))
	fmt.Printf("g. Standar deviasi: %.2f\n", stdDev)

	fmt.Print("Masukkan nilai untuk mencari frekuensi: ")
	fmt.Scan(&value)
	frequency := 0
	for _, num := range array {
		if num == value {
			frequency++
		}
	}
	fmt.Printf("h. Frekuensi %d: %d kali\n", value, frequency)
}
