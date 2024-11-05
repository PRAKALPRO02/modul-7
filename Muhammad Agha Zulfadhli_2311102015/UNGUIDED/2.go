package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("a. Tampilkan elemen-elemen array")
	fmt.Println("b. Tampilkan elemen-elemen array ganjil")
	fmt.Println("c. Tampilkan elemen-elemen array genap")
	fmt.Println("d. Tampilkan elemen-elemen array kelipatan \"x\"")
	fmt.Println("e. Hapus indeks ke-\"x\" dari array")
	fmt.Println("f. Tampilkan rata-rata elemen array")
	fmt.Println("g. Standar deviasi dari array")
	fmt.Println("h. Tampilkan frekuensi \"x\" dari array")
	fmt.Println("____________________________________________________________")
	fmt.Print("Pilihan : ")

	in := ""
	fmt.Scan(&in)
	fmt.Println()

	switch in {
	case "a": //? SHOW ELEMENTS
		print_arr(arr)
	case "b":
		for i := 0; i < len(arr); i++ {
			if arr[i]%2 != 0 {
				fmt.Print(arr[i], " ")
			}
		}
	case "c": //? SHOW GANJIL
		for i := 0; i < len(arr); i++ {
			if arr[i]%2 == 0 {
				fmt.Print(arr[i], " ")
			}
		}
	case "d": //? SHOW GENAP
		num := 0
		fmt.Print("Nilai yang ingin dicari : ")
		fmt.Scan(&num)
		for i, value := range arr {
			if value == num {
				fmt.Print("Nilai yang dicari ada di index ke-", i+1)
			}
		}
	case "e": //? DELETE ELEMENT
		cut := 0
		fmt.Print("Index yang ingin dihapus : ")
		fmt.Scan(&cut)
		if cut < 0 {
			fmt.Println("Array Mulai dari 0")
		} else if cut > len(arr) {
			fmt.Println("Index diluar array")
			break
		}
		fmt.Print("Sebelum hapus : ")
		print_arr(arr)
		arr = append(arr[:cut], arr[cut+1:]...)
		fmt.Print("Selesai hapus : ")
		print_arr(arr)
	case "f": //? RATA RATA
		sum := 0
		for _, value := range arr {
			sum += value
		}
		fmt.Println(float64(sum) / float64(len(arr)))
	case "g": //? STANDARD DEVIATION
		fmt.Print("Array : ")
		print_arr(arr)
		fmt.Println(standardDeviation(arr))
	case "h": //? FREKUENSI
		var num, all int
		fmt.Print("Nilai yang ingin dicari : ")
		fmt.Scan(&num)
		for _, value := range arr {
			if value == num {
				all++
			}
		}
		fmt.Println(all)
	}
}

func print_arr(arr []int) {
	for _, value := range arr {
		fmt.Print(value, " ")
	}
	fmt.Println()
}

func standardDeviation(arr []int) float64 {
	var sum, mean, variance float64
	n := float64(len(arr))

	for _, value := range arr {
		sum += float64(value)
	}
	mean = sum / n

	for _, value := range arr {
		variance += math.Pow(float64(value)-mean, 2)
	}
	return math.Sqrt(variance / n)
}
