package main

import (
	"fmt"
	"math"
)

func tampilkanArray225(arr []int) {
	for i, val := range arr {
		fmt.Printf("Indeks %d: %d\n", i, val)
	}
}

func tampilkanIndeksGanjil225(arr []int) {
	for i := 1; i < len(arr); i += 2 {
		fmt.Printf("Indeks %d (ganjil): %d\n", i, arr[i])
	}
}

func tampilkanIndeksGenap225(arr []int) {
	for i := 0; i < len(arr); i += 2 {
		fmt.Printf("Indeks %d (genap): %d\n", i, arr[i])
	}
}

func tampilkanIndeksKelipatan225(arr []int, x int) {
	for i := 0; i < len(arr); i += x {
		fmt.Printf("Indeks %d (kelipatan %d): %d\n", i, x, arr[i])
	}
}

func hapusElemen225(arr []int, indeks int) []int {
	return append(arr[:indeks], arr[indeks+1:]...)
}

func rataRata225(arr []int) float64 {
	total := 0
	for _, val := range arr {
		total += val
	}
	return float64(total) / float64(len(arr))
}

func standarDeviasi225(arr []int) float64 {
	mean := rataRata225(arr)
	var sum float64
	for _, val := range arr {
		sum += math.Pow(float64(val)-mean, 2)
	}
	return math.Sqrt(sum / float64(len(arr)))
}

func hitungFrekuensi225(arr []int, target int) int {
	frekuensi := 0
	for _, val := range arr {
		if val == target {
			frekuensi++
		}
	}
	return frekuensi
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah elemen dalam array: ")
	fmt.Scan(&n)

	arr := make([]int, n)

	fmt.Println("Masukkan elemen-elemen array:")
	for i := 0; i < n; i++ {
		fmt.Printf("Elemen ke-%d: ", i)
		fmt.Scan(&arr[i])
	}

	fmt.Println("Seluruh isi array:")
	tampilkanArray225(arr)

	fmt.Println("Elemen dengan indeks ganjil:")
	tampilkanIndeksGanjil225(arr)

	fmt.Println("Elemen dengan indeks genap:")
	tampilkanIndeksGenap225(arr)

	var x int
	fmt.Print("Masukkan bilangan x untuk kelipatan indeks: ")
	fmt.Scan(&x)
	fmt.Println("Elemen dengan indeks kelipatan", x, ":")
	tampilkanIndeksKelipatan225(arr, x)

	var indeksHapus int
	fmt.Print("Masukkan indeks yang ingin dihapus: ")
	fmt.Scan(&indeksHapus)
	if indeksHapus >= 0 && indeksHapus < len(arr) {
		arr = hapusElemen225(arr, indeksHapus)
		fmt.Println("Array setelah elemen dihapus:")
		tampilkanArray225(arr)
	} else {
		fmt.Println("Indeks tidak valid!")
	}

	fmt.Printf("Rata-rata: %.2f\n", rataRata225(arr))

	fmt.Printf("Standar deviasi: %.2f\n", standarDeviasi225(arr))

	var target int
	fmt.Print("Masukkan bilangan untuk menghitung frekuensinya: ")
	fmt.Scan(&target)
	fmt.Printf("Frekuensi bilangan %d: %d kali\n", target, hitungFrekuensi225(arr, target))
}
