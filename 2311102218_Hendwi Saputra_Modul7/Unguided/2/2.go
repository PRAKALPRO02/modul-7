package main

import (
	"fmt"
	"math"
)

func main() {
	var N int

	fmt.Print("Masukkan jumlah elemn array: ")
	fmt.Scan(&N)

	arr := make([]int, N)

	fmt.Println("Masukkan nilai-nilai array:")
	for i := 0; i < N; i++ {
		fmt.Printf("Elemen ke-%d: ", i+1)
		fmt.Scan(&arr[i])
	}

	fmt.Println("\nSeluruh isi array:")
	for _, v := range arr {
		fmt.Print(v, " ")
	}
	fmt.Println()

	fmt.Println("\nElemen-elemen array dengan indeks ganjil:")
	for i := 1; i < N; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	fmt.Println("\nElemen-elemen array dengan indeks gendap:")
	for i := 0; i < N; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	var x int
	fmt.Print("\nMasukkan nilai x: ")
	fmt.Scan(&x)
	fmt.Printf("Elemen-elemen array dengan indeks kelipatan %d: \n", x)
	for i := 0; i < N; i++ {
		if i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	var hapusElemen int
	fmt.Print("\nMasukkan indeks yang ingin dihapus: ")
	fmt.Scan(&hapusElemen)
	if hapusElemen >= 0 && hapusElemen < len(arr) {
		arr = append(arr[:hapusElemen], arr[hapusElemen+1:]...)
		fmt.Println("Array setelah penghapusan:")
		for _, v := range arr {
			fmt.Print(v, " ")
		}
		fmt.Println()
	} else {
		fmt.Println("Indeks tidak valid")
	}

	var sum float64
	for _, v := range arr {
		sum += float64(v)
	}
	rataRata := sum / float64(len(arr))
	fmt.Printf("\nRata-rata: %.2f\n", rataRata)

	var sumKuadrat float64
	for _, v := range arr {
		sumKuadrat += float64(v * v)
	}
	varian := (sumKuadrat / float64(len(arr))) - math.Pow(rataRata, 2)
	standarDeviasi := math.Sqrt(varian)
	fmt.Printf("Standar deviasi: %.2f\n", standarDeviasi)

	var bilangan int
	fmt.Print("\nMasukkan bilangan yang ingin dicari frekuensinya: ")
	fmt.Scan(&bilangan)
	var frekuensi int
	for _, v := range arr {
		if v == bilangan {
			frekuensi++
		}
	}
	fmt.Printf("Frekuensi %d: %d\n", bilangan, frekuensi)
}
