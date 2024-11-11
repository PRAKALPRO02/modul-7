package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Masukkan jumlah elemen dalam array: ")
	fmt.Scan(&n)

	array := make([]int, n)
	fmt.Println("Masukkan elemen-elemen array:")
	for i := 0; i < n; i++ {
		fmt.Printf("Elemen ke-%d: ", i)
		fmt.Scan(&array[i])
	}

	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Tampilkan seluruh isi array")
		fmt.Println("2. Tampilkan elemen dengan indeks ganjil")
		fmt.Println("3. Tampilkan elemen dengan indeks genap")
		fmt.Println("4. Tampilkan elemen dengan indeks kelipatan X")
		fmt.Println("5. Hapus elemen pada indeks tertentu")
		fmt.Println("6. Tampilkan rata-rata dari elemen array")
		fmt.Println("7. Tampilkan standar deviasi dari elemen array")
		fmt.Println("8. Keluar")
		fmt.Print("Pilih opsi: ")

		var pilih int
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			fmt.Println("Isi array:", array)

		case 2:
			fmt.Println("Elemen dengan indeks ganjil:")
			for i := 1; i < len(array); i += 2 {
				fmt.Printf("Index %d: %d\n", i, array[i])
			}

		case 3:
			fmt.Println("Elemen dengan indeks genap:")
			for i := 0; i < len(array); i += 2 {
				fmt.Printf("Index %d: %d\n", i, array[i])
			}

		case 4:
			var x int
			fmt.Print("Masukkan nilai X: ")
			fmt.Scan(&x)
			fmt.Printf("Elemen dengan indeks kelipatan %d:\n", x)
			for i := 0; i < len(array); i++ {
				if i%x == 0 {
					fmt.Printf("Index %d: %d\n", i, array[i])
				}
			}

		case 5:
			var index int
			fmt.Print("Masukkan indeks yang ingin dihapus: ")
			fmt.Scan(&index)
			if index >= 0 && index < len(array) {
				array = append(array[:index], array[index+1:]...)
				fmt.Println("Array setelah penghapusan:", array)
			} else {
				fmt.Println("Indeks tidak valid.")
			}

		case 6:
			sum := 0
			for _, val := range array {
				sum += val
			}
			rerata := float64(sum) / float64(len(array))
			fmt.Printf("Rata-rata elemen array: %.2f\n", rerata)

		case 7:
			sum := 0
			for _, val := range array {
				sum += val
			}
			mean := float64(sum) / float64(len(array))

			variasiTambah := 0.0
			for _, val := range array {
				diff := float64(val) - mean
				variasiTambah += diff * diff
			}
			variasi := variasiTambah / float64(len(array))
			deviasi := math.Sqrt(variasi)
			fmt.Printf("Standar deviasi elemen array: %.2f\n", deviasi)

		case 8:
			fmt.Println("Keluar dari program.")
			return

		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
} 