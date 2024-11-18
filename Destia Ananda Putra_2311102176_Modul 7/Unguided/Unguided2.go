package main

// Destia Ananda Putra
// 2311102176

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Masukkan jumlah elemen array (N): ")
	fmt.Scan(&n)

	arr := make([]int, n)
	fmt.Println("Masukkan elemen-elemen array:")
	for i := 0; i < n; i++ {
		fmt.Printf("Elemen ke-%d: ", i)
		fmt.Scan(&arr[i])
	}

	fmt.Println("a. Isi seluruh array:", arr)

	fmt.Print("b. Elemen dengan indeks ganjil: ")
	for i := 1; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	fmt.Print("c. Elemen dengan indeks genap: ")
	for i := 0; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	var x int
	fmt.Print("d. Masukkan angka x untuk melihat elemen dengan indeks kelipatan x: ")
	fmt.Scan(&x)
	fmt.Printf("Elemen dengan indeks kelipatan %d: ", x)
	for i := x; i < n; i += x {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	var index int
	fmt.Print("e. Masukkan indeks elemen yang ingin dihapus: ")
	fmt.Scan(&index)
	if index >= 0 && index < n {
		arr = append(arr[:index], arr[index+1:]...)
		fmt.Println("Array setelah elemen dihapus:", arr)
	} else {
		fmt.Println("Indeks tidak valid.")
	}

	sum := 0
	for _, value := range arr {
		sum += value
	}
	rataRata := float64(sum) / float64(len(arr))
	fmt.Printf("f. Rata-rata elemen array: %.2f\n", rataRata)

	var deviasi float64
	for _, value := range arr {
		deviasi += math.Pow(float64(value)-rataRata, 2)
	}
	deviasi /= float64(len(arr))
	stdDev := math.Sqrt(deviasi)
	fmt.Printf("g. Standar deviasi elemen array: %.2f\n", stdDev)

	frekuensi := make(map[int]int)
	for _, value := range arr {
		frekuensi[value]++
	}
	fmt.Println("h. Frekuensi kemunculan setiap elemen:")
	for key, freq := range frekuensi {
		fmt.Printf("Angka %d muncul %d kali\n", key, freq)
	}
}
