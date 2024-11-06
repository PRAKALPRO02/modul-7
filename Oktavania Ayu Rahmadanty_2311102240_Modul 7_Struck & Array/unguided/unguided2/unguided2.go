package main

import (
	"fmt"
	"math"
)

func isiArray(arr *[]int, n int) {
	*arr = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan elemen ke-%d: ", i)
		fmt.Scan(&(*arr)[i])
	}
}

func tampilkanArray(arr []int) {
	fmt.Println("Isi array:", arr)
}

func tampilkanIndeksGanjil(arr []int) {
	fmt.Print("Elemen di indeks genap: ")
	for i := 1; i < len(arr); i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}

func tampilkanIndeksGenap(arr []int) {
	fmt.Print("Elemen di indeks ganjil: ")
	for i := 0; i < len(arr); i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}

func tampilkanIndeksKelipatan(arr []int, x int) {
	fmt.Printf("Elemen di indeks kelipatan %d: ", x)
	for i := x; i < len(arr); i += x {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}

func hapusElemen(arr *[]int, indeks int) {
	*arr = append((*arr)[:indeks], (*arr)[indeks+1:]...)
	fmt.Println("Isi array setelah penghapusan:", *arr)
}

func hitungRataRata(arr []int) float64 {
	total := 0
	for _, v := range arr {
		total += v
	}
	return float64(total) / float64(len(arr))
}

func hitungStandarDeviasi(arr []int) float64 {
	rata := hitungRataRata(arr)
	var total float64
	for _, v := range arr {
		total += math.Pow(float64(v)-rata, 2)
	}
	return math.Sqrt(total / float64(len(arr)))
}

func hitungFrekuensi(arr []int, bilangan int) int {
	count := 0
	for _, v := range arr {
		if v == bilangan {
			count++
		}
	}
	return count
}

func main() {
	var n, x, indeks, bilangan int

	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scan(&n)

	arr := []int{}
	isiArray(&arr, n)

	fmt.Println("\nMenu:")
	fmt.Println("1. Tampilkan seluruh elemen array")
	fmt.Println("2. Tampilkan elemen dengan indeks genap")
	fmt.Println("3. Tampilkan elemen dengan indeks ganjil")
	fmt.Println("4. Tampilkan elemen dengan indeks kelipatan x")
	fmt.Println("5. Hapus elemen pada indeks tertentu")
	fmt.Println("6. Tampilkan rata-rata elemen")
	fmt.Println("7. Tampilkan standar deviasi elemen")
	fmt.Println("8. Tampilkan frekuensi dari suatu bilangan")
	fmt.Println("9. Keluar")

	for {
		var pilihan int
		fmt.Print("\nPilih menu: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tampilkanArray(arr)
		case 2:
			tampilkanIndeksGanjil(arr)
		case 3:
			tampilkanIndeksGenap(arr)
		case 4:
			fmt.Print("Masukkan nilai x: ")
			fmt.Scan(&x)
			if x > 0 {
				tampilkanIndeksKelipatan(arr, x)
			} else {
				fmt.Println("Nilai x harus lebih besar dari 0.")
			}
		case 5:
			fmt.Print("Masukkan indeks yang ingin dihapus: ")
			fmt.Scan(&indeks)
			if indeks >= 0 && indeks < len(arr) {
				hapusElemen(&arr, indeks)
			} else {
				fmt.Println("Indeks tidak valid.")
			}
		case 6:
			fmt.Printf("Rata-rata elemen: %.2f\n", hitungRataRata(arr))
		case 7:
			fmt.Printf("Standar deviasi elemen: %.2f\n", hitungStandarDeviasi(arr))
		case 8:
			fmt.Print("Masukkan bilangan untuk dihitung frekuensinya: ")
			fmt.Scan(&bilangan)
			fmt.Printf("Frekuensi %d: %d\n", bilangan, hitungFrekuensi(arr, bilangan))
		case 9:
			fmt.Println("Keluar program.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
