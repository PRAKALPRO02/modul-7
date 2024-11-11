package main

import (
	"fmt"
	"math"
)

func main() {
	var n, idx213, x213 int
	fmt.Print("Masukkan jumlah elemen array (N): ")
	fmt.Scan(&n)

	array := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan elemen ke-%d: ", i)
		fmt.Scan(&array[i])
	}

	fmt.Println("a. Menampilkan keseluruhan isi dari array:")
	printarray(array)

	fmt.Println("b. Menampilkan elemen-elemen array yang ganjil saja:")
	printganjil(array)

	fmt.Println("c. Menampilkan elemen-elemen array yang genap saja:")
	printgenap(array)

	fmt.Print("Masukkan bilangan untuk kelipatan (x213): ")
	fmt.Scan(&x213)
	fmt.Printf("d. Menampilkan elemen-elemen array dengan indeks kelipatan %d:\n", x213)
	kelipatan2(array, x213)

	fmt.Print("Masukkan indeks yang ingin dihapus: ")
	fmt.Scan(&idx213)
	array = hapus(array, idx213)
	fmt.Println("e. Array setelah penghapusan:")
	printarray(array)

	fmt.Println("f. Rata-rata dari bilangan di dalam array:")
	fmt.Println(average(array))

	fmt.Println("g. Standar deviasi dari bilangan di dalam array:")
	fmt.Println(rerata(array))

	fmt.Println("h. Menampilkan nilai terbesar dan terkecil dalam array:")
	maxxxx(array)
}

func printarray(arr []int) {
	for _, v := range arr {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func printganjil(arr []int) {
	for _, v := range arr {
		if v%2 != 0 {
			fmt.Print(v, " ")
		}
	}
	fmt.Println()
}

func printgenap(arr []int) {
	for _, v := range arr {
		if v%2 == 0 {
			fmt.Print(v, " ")
		}
	}
	fmt.Println()
}

func kelipatan2(arr []int, x213 int) {
	for i, v := range arr {
		if i%x213 == 0 {
			fmt.Print(v, " ")
		}
	}
	fmt.Println()
}

func hapus(arr []int, idx213 int) []int {
	return append(arr[:idx213], arr[idx213+1:]...)
}

func average(arr []int) float64 {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return float64(sum) / float64(len(arr))
}

func rerata(arr []int) float64 {
	mean := average(arr)
	var variance float64
	for _, v := range arr {
		variance += math.Pow(float64(v)-mean, 2)
	}
	variance /= float64(len(arr))
	return math.Sqrt(variance)
}

func maxxxx(arr []int) {
	if len(arr) == 0 {
		fmt.Println("Array kosong.")
		return
	}
	min, max213 := arr[0], arr[0]
	for _, v := range arr {
		if v < min {
			min = v
		}
		if v > max213 {
			max213 = v
		}
	}
	fmt.Println("Nilai terkecil:", min)
	fmt.Println("Nilai terbesar:", max213)
}
