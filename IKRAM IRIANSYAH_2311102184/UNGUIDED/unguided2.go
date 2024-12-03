package main

import (
	"fmt"
	"math"
)

type Array []int

func show_array(arr []int) {
	fmt.Println("ARRAY :", arr)
}

func index_ganjil(arr []int) {
	fmt.Print("Indeks Ganjil : ")
	for i := 1; i < len(arr); i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}

func index_genap(arr []int) {
	fmt.Print("Indeks Genap : ")
	for i := 0; i < len(arr); i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}

func kelipatan_index(arr []int, x int) {
	fmt.Printf("Indeks Kelipatan %d : ", x)
	for i := 0; i < len(arr); i += x {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}

func hapus_index(arr []int, index int) []int {
	if index < 0 || index >= len(arr) {
		fmt.Println("Indeks Tidak Valid")
		return arr
	}
	fmt.Printf("Elemen di Indeks %d (nilai : %d) dihapus!\n", index, arr[index])
	return append(arr[:index], arr[index+1:]...)
}

func rata_rata(arr []int) float64 {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return float64(sum) / float64(len(arr))
}

func deviasi(arr []int) float64 {
	mean := rata_rata(arr)
	var sum float64
	for _, v := range arr {
		sum += math.Pow(float64(v)-mean, 2)
	}
	return math.Sqrt(sum / float64(len(arr)))
}

func frekuensi(arr []int, nilai int) int {
	count := 0
	for _, v := range arr {
		if v == nilai {
			count++
		}
	}
	return count
}

func main() {
	var n, x, index, nilai int

	fmt.Print("Masukkan jumlah elemen array : ")
	fmt.Scan(&n)

	arr := make(Array, n)
	fmt.Print("Masukkan elemen array : ")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	show_array(arr)
	index_ganjil(arr)
	index_genap(arr)

	fmt.Print("\nMasukkan nilai x untuk kelipatan indeks : ")
	fmt.Scan(&x)
	kelipatan_index(arr, x)

	fmt.Print("\nMasukkan indeks untuk menghapus elemen : ")
	fmt.Scan(&index)
	arr = hapus_index(arr, index)
	show_array(arr)

	fmt.Printf("\nRATA RATA : %.2f\n", rata_rata(arr))

	fmt.Printf("DEVIASI : %.2f\n", deviasi(arr))

	fmt.Print("\nMasukkan nilai untuk menghitung frekuensi : ")
	fmt.Scan(&nilai)
	fmt.Printf("FREKUENSI %d : %d\n", nilai, frekuensi(arr, nilai))
}
