package main

import (
	"fmt"
	"math"
)

func main() {
	var N, x, indexHapus, target int
	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scan(&N)

	arr := make([]int, N)
	for i := range arr {
		fmt.Printf("Masukkan elemen ke-%d: ", i)
		fmt.Scan(&arr[i])
	}

	fmt.Println("Isi array:", arr)

	fmt.Print("Elemen dengan indeks ganjil: ")
	for i := 1; i < N; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	fmt.Print("Elemen dengan indeks genap: ")
	for i := 0; i < N; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	fmt.Print("Masukkan bilangan x untuk indeks kelipatan: ")
	fmt.Scan(&x)
	fmt.Printf("Elemen dengan indeks kelipatan %d: ", x)
	for i := 0; i < N; i += x {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	fmt.Print("Masukkan indeks yang ingin dihapus: ")
	fmt.Scan(&indexHapus)
	arr = append(arr[:indexHapus], arr[indexHapus+1:]...)
	fmt.Println("Isi array setelah penghapusan:", arr)

	var sum int
	for _, v := range arr {
		sum += v
	}
	average := float64(sum) / float64(len(arr))
	fmt.Printf("Rata-rata: %.2f\n", average)

	var varianceSum float64
	for _, v := range arr {
		varianceSum += math.Pow(float64(v)-average, 2)
	}
	stdDev := math.Sqrt(varianceSum / float64(len(arr)))
	fmt.Printf("Standar deviasi: %.2f\n", stdDev)

	fmt.Print("Masukkan bilangan untuk frekuensi: ")
	fmt.Scan(&target)
	count := 0
	for _, v := range arr {
		if v == target {
			count++
		}
	}
	fmt.Printf("Frekuensi dari bilangan %d: %d kali\n", target, count)
}
