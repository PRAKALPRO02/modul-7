// 2311102162
package main

import (
	"fmt"
	"math"
)

func main() {
	var N int
	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scan(&N)

	array := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Printf("Masukkan elemen ke-%d: ", i)
		fmt.Scan(&array[i])
	}

	fmt.Println("\nIsi array:")
	fmt.Println(array)

	fmt.Println("\nElemen dengan indeks ganjil:")
	for i := 1; i < N; i += 2 {
		fmt.Print(array[i], " ")
	}

	fmt.Println("\n\nElemen dengan indeks genap:")
	for i := 0; i < N; i += 2 {
		fmt.Print(array[i], " ")
	}

	var x int
	fmt.Println("\n\nElemen dengan kelipatan x:")
	fmt.Print("Masukkan bilangan x: ")
	fmt.Scan(&x)
	fmt.Println("Elemen dengan indeks kelipatan", x, ":")
	for i := 0; i < N; i++ {
		if i%x == 0 {
			fmt.Print(array[i], " ")
		}
	}

	var indeksHapus int
	fmt.Print("\n\nMasukkan indeks yang akan dihapus: ")
	fmt.Scan(&indeksHapus)
	array = append(array[:indeksHapus], array[indeksHapus+1:]...)
	fmt.Println("Array setelah dihapus elemen pada indeks", indeksHapus, ":")
	fmt.Println(array)

	var total int
	for _, value := range array {
		total += value
	}
	rataRata := float64(total) / float64(len(array))
	fmt.Printf("\nRata-rata: %.2f\n", rataRata)

	var variance float64
	for _, value := range array {
		variance += math.Pow(float64(value)-rataRata, 2)
	}
	variance /= float64(len(array))
	standarDeviasi := math.Sqrt(variance)
	fmt.Printf("Standar deviasi: %.2f\n", standarDeviasi)

	var bilangan int
	fmt.Print("\nMasukkan bilangan yang frekuensinya ingin dihitung: ")
	fmt.Scan(&bilangan)
	frekuensi := 0
	for _, value := range array {
		if value == bilangan {
			frekuensi++
		}
	}
	fmt.Printf("Frekuensi bilangan %d: %d\n", bilangan, frekuensi)
}
