package main

import (
	"fmt"
	"math"
)

func faktorBilangan234(bilangan234, i234 int) {
	if i234 > bilangan234 {
		return
	}
	if bilangan234%i234 == 0 {
		fmt.Print(i234, " ")
	}
	faktorBilangan234(bilangan234, i234+1)
}

func main() {
	var N234 int
	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scan(&N234)

	array234 := make([]int, N234)

	for i234 := 0; i234 < N234; i234++ {
		fmt.Printf("Masukkan elemen ke-%d: ", i234)
		fmt.Scan(&array234[i234])
	}

	fmt.Println("\nIsi array:")
	fmt.Println(array234)

	fmt.Println("\nElemen dengan indeks ganjil:")
	for i234 := 1; i234 < N234; i234 += 2 {
		fmt.Print(array234[i234], " ")
	}

	fmt.Println("\n\nElemen dengan indeks genap:")
	for i234 := 0; i234 < N234; i234 += 2 {
		fmt.Print(array234[i234], " ")
	}

	var x234 int
	fmt.Print("\n\nMasukkan bilangan x: ")
	fmt.Scan(&x234)
	fmt.Println("Elemen dengan indeks kelipatan", x234, ":")
	for i234 := 0; i234 < N234; i234++ {
		if i234%x234 == 0 {
			fmt.Print(array234[i234], " ")
		}
	}

	var indeksHapus234 int
	fmt.Print("\n\nMasukkan indeks yang akan dihapus: ")
	fmt.Scan(&indeksHapus234)
	if indeksHapus234 >= 0 && indeksHapus234 < len(array234) {
		array234 = append(array234[:indeksHapus234], array234[indeksHapus234+1:]...)
		fmt.Println("Array setelah dihapus elemen pada indeks", indeksHapus234, ":")
		fmt.Println(array234)
	} else {
		fmt.Println("Indeks tidak valid.")
	}

	var total234 int
	for _, value234 := range array234 {
		total234 += value234
	}
	if len(array234) > 0 {
		rataRata234 := float64(total234) / float64(len(array234))
		fmt.Printf("\nRata-rata: %.2f\n", rataRata234)

		var variance234 float64
		for _, value234 := range array234 {
			variance234 += math.Pow(float64(value234)-rataRata234, 2)
		}
		variance234 /= float64(len(array234))
		standarDeviasi234 := math.Sqrt(variance234)
		fmt.Printf("Standar deviasi: %.2f\n", standarDeviasi234)
	} else {
		fmt.Println("Array kosong, tidak dapat menghitung rata-rata dan standar deviasi.")
	}

	var bilangan234 int
	fmt.Print("\nMasukkan bilangan yang frekuensinya ingin dihitung: ")
	fmt.Scan(&bilangan234)
	frekuensi234 := 0
	for _, value234 := range array234 {
		if value234 == bilangan234 {
			frekuensi234++
		}
	}
	fmt.Printf("Frekuensi bilangan %d: %d\n", bilangan234, frekuensi234)

	fmt.Print("\nMasukkan bilangan untuk mencari faktornya: ")
	var inputBilangan234 int
	fmt.Scan(&inputBilangan234)
	fmt.Println("Faktor dari", inputBilangan234, "adalah:")
	faktorBilangan234(inputBilangan234, 1)
	fmt.Println()
}
