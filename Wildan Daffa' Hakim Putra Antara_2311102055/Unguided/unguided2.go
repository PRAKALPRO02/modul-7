package main

import (
	"fmt"
	"math"
)

func showMenu() {
	fmt.Println("1. menampilkan seluruh isi array")
	fmt.Println("2. Menampilkan elemen-elemen array dengan indeks ganjil saja.")
	fmt.Println("3. Menampilkan elemen-elemen array dengan indeks genap saja (asumsi indek ke-0 adalah genap).")
	fmt.Println("4. Menampilkan elemen-elemen array dengan indeks kelipatan bilangan x")
	fmt.Println("5. Menghapus elemen array pada indeks tertentu.")
	fmt.Println("6. Menampilkan rata-rata dari bilangan yang ada di dalam array.")
	fmt.Println("7. Menampilkan standar deviasi atau simpangan baku")
	fmt.Println("8. Menampilkan frekuensi dari suatu bilangan")
	fmt.Println("9. Exit")
	fmt.Print("Pilih menu : ")
}

func showData(arr []int, skip int, expect int, skipZero bool) {
	for index, data := range arr {
		if skipZero && index == 0 {
			continue
		}
		if index%skip == expect {
			fmt.Println("Index ke - ", index, " : ", data)
		}
	}
}

func reRata(arr []int) float64 {
	var sum int
	for _, data := range arr {
		sum += data
	}
	return float64(sum) / float64(len(arr))
}

func simpBaku(arr []int) float64 {
	var sum float64
	for _, data := range arr {
		sum += math.Pow(float64(data)-reRata(arr), 2)
	}
	return math.Sqrt(float64(sum) / float64(len(arr)))
}

func frek(arr []int, bil int) int {
	var count int
	for _, data := range arr {
		if data == bil {
			count++
		}
	}
	return count
}

func removeIndex(index int, arr *[]int) {
	if len(*arr)-1 < index {
		fmt.Println("index out of range")
	} else {
		fmt.Println("item", (*arr)[index], "pada index", index, "telah dihapus\nisi Sekarang : ")
		*arr = append((*arr)[:index], (*arr)[index+1:]...)
		showData(*arr, 1, 0, false)
	}
}

func main() {
	var arraySize, input int
	fmt.Print("Masukkan Kapasitas array : ")
	fmt.Scan(&arraySize)
	var newArr []int = make([]int, arraySize)
	for index, _ := range newArr {
		fmt.Print("Masukkan index ke -", index, " : ")
		fmt.Scan(&newArr[index])
	}
	for input != 9 {
		showMenu()
		fmt.Scan(&input)
		switch input {
		case 1:
			showData(newArr, 1, 0, false)
		case 2:
			showData(newArr, 2, 1, false)
		case 3:
			showData(newArr, 2, 0, false)
		case 4:
			fmt.Print("Masukkan Kelipatan : ")
			fmt.Scan(&input)
			showData(newArr, input, 0, true)
		case 5:
			fmt.Print("Masukkan index yang dihapus : ")
			fmt.Scan(&input)
			removeIndex(input, &newArr)
		case 6:
			fmt.Println("Rerata dari isi array adalah :", reRata(newArr))
		case 7:
			fmt.Println("Deviasi dari isi array adalah :", simpBaku(newArr))
		case 8:
			fmt.Print("Masukkan bilangan yang akan dicek : ")
			fmt.Scan(&input)
			fmt.Println("Frekuansi dari", input, "pada array adalah :", frek(newArr, input))
		default:
			fmt.Println("Opsi salah")
		}
	}
}
