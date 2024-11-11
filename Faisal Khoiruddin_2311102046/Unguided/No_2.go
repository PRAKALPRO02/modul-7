package main

import (
	"fmt"
	"math"
)

func isiArray(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan nilai untuk elemen %d: ", i)
		fmt.Scan(&arr[i])
	}
	return arr
}

func displayArray(arr []int) {
	fmt.Println("Isi aray:", arr)
}

func displayIndeksGanjil(arr []int) {
	fmt.Print("Elemen dengan indeks ganji: ")
	for i := 1; i < len(arr); i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}

func displayIndeksGenap(arr []int) {
	fmt.Print("Elemen dengan indeks genap: ")
	for i := 0; i < len(arr); i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}

func displayIndeksKelipatan(arr []int, x int) {
	fmt.Printf("Elemen dengan indeks kelipatan %d: ", x)
	for i := 0; i < len(arr); i++ {
		if i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()
}

func deleteElemen(arr []int, indeks int) []int {
	return append(arr[:indeks], arr[indeks+1:]...)
}

func hitungRataRata(arr []int) float64 {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return float64(sum) / float64(len(arr))
}

func hitungStandarDeviasi(arr []int) float64 {
	mean := hitungRataRata(arr)
	sum := 0.0
	for _, val := range arr {
		sum += math.Pow(float64(val)-mean, 2)
	}
	return math.Sqrt(sum / float64(len(arr)))
}

func hitungFrekuensi(arr []int, number int) int {
	count := 0
	for _, val := range arr {
		if val == number {
			count++
		}
	}
	return count
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah elemen: ")
	fmt.Scan(&n)

	arr := isiArray(n)

	var pilih int
	for {
		fmt.Println("\n==================================================")
		fmt.Println("|                Program pada Array	         |")
		fmt.Println("==================================================")
		fmt.Println("| 1. | Tampilkan array                           |")
		fmt.Println("| 2. | Tampilkan elemen dengan indeks ganjil     |")
		fmt.Println("| 3. | Tampilkan elemen dengan indeks genap      |")
		fmt.Println("| 4. | Tampilkan elemen dengan indeks kelipatan x|")
		fmt.Println("| 5. | Hapus elemen pada indeks tertentu         |")
		fmt.Println("| 6. | Hitung rata-rata elemen array             |")
		fmt.Println("| 7. | Hitung standar deviasi elemen array       |")
		fmt.Println("| 8. | Hitung frekuensi elemen tertentu          |")
		fmt.Println("| 9. | keluar                                    |")
		fmt.Println("==================================================")

		fmt.Print("Pilih operasi: ")
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			displayArray(arr)
		case 2:
			displayIndeksGanjil(arr)
		case 3:
			displayIndeksGenap(arr)
		case 4:
			var x int
			fmt.Print("Masukkan nilai x: ")
			fmt.Scan(&x)
			displayIndeksKelipatan(arr, x)
		case 5:
			var indeks int
			fmt.Print("Masukkan indeks yang dihapus: ")
			fmt.Scan(&indeks)
			arr = deleteElemen(arr, indeks)
			displayArray(arr)
		case 6:
			RataRata := hitungRataRata(arr)
			fmt.Println("Rata-Rata pada array: ", RataRata)
		case 7:
			standarDeviasi := hitungStandarDeviasi(arr)
			fmt.Println("Standar deviasi dari elemen array: ", standarDeviasi)
		case 8:
			var number int
			fmt.Print("Masukkan bilangan untuk menghitung frekuensi: ")
			fmt.Scan(&number)
			frekuensi := hitungFrekuensi(arr, number)
			fmt.Printf("Frekuensi %d dalam array: %d\n", number, frekuensi)
		case 9:
			fmt.Println("Keluar dari program.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
