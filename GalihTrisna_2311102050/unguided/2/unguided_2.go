package main

import (
	"fmt"
	"math"
)

func genap(arr *[]int) {
	fmt.Print("[")
	for i := 0; i < len(*arr); i++ {
		if i%2 == 0 {
			fmt.Printf("%v, ", (*arr)[i])
		}
	}
	fmt.Print("]\n")
}

func ganjil(arr *[]int) {
	fmt.Print("[")
	for i := 0; i < len(*arr); i++ {
		if i%2 != 0 {
			fmt.Printf("%v, ", (*arr)[i])
		}
	}
	fmt.Print("]\n")
}

func kelipatan(arr *[]int) {
	var x int
	fmt.Print("Masukkan nilai x: ")
	fmt.Scanln(&x)
	fmt.Print("[")
	for i := 0; i < len(*arr); i++ {
		if (*arr)[i]%x == 0 {
			fmt.Printf("%v, ", (*arr)[i])
		}
	}
	fmt.Print("]\n")
}

func hapus(arr *[]int) {
	var indeks int
	fmt.Print("Masukkan indeks: ")
	fmt.Scanln(&indeks)
	*arr = append((*arr)[:indeks], (*arr)[indeks+1:]...)
	fmt.Printf("%v\n", arr)
}

func rata2(arr *[]int) {
	var total int
	for i := 0; i < len(*arr); i++ {
		total += (*arr)[i]
	}
	fmt.Printf("Rata - rata: %v\n", total/len(*arr))
}

func simpangan_baku(arr *[]int) {
	var total, rerata, simpangan_baku float64
	for i := 0; i < len(*arr); i++ {
		total += float64((*arr)[i])
	}
	rerata = total / float64(len(*arr))
	for i := 0; i < len(*arr); i++ {
		simpangan_baku += math.Pow(float64((*arr)[i])-rerata, 2)
	}
	simpangan_baku = math.Sqrt(simpangan_baku / float64(len((*arr))))
	fmt.Printf("Simpangan baku dari array adalah %v\n", simpangan_baku)
}

func frekuensi(arr *[]int) {
	var x, total int
	fmt.Print("Masukkan nilai x: ")
	fmt.Scanln(&x)
	for i := 0; i < len(*arr); i++ {
		if (*arr)[i] == x {
			total++
		}
	}
	fmt.Printf("Frekuensi %v dalam array adalah %v\n", x, total)
}

func main() {
	var arrSize, menu int
	fmt.Print("Masukkan ukuran array: ")
	fmt.Scanln(&arrSize)
	arr := make([]int, arrSize)
	for i := 0; i < len(arr); i++ {
		fmt.Printf("Masukkan nilai elemen ke - %v: ", i+1)
		fmt.Scanln(&arr[i])
	}
	for true {
		fmt.Printf("\n[PROGRAM MENU]\n")
		fmt.Printf("1. Tampilkan semua elemen\n")
		fmt.Printf("2. Tampilkan elemen indeks ganjil\n")
		fmt.Printf("3. Tampilkan elemen indeks genap\n")
		fmt.Printf("4. Tampilkan elemen kelipatan x\n")
		fmt.Printf("5. Hapus elemen\n")
		fmt.Printf("6. Tampilkan rata - rata semua elemen\n")
		fmt.Printf("7. Tampilkan simpangan baku\n")
		fmt.Printf("8. Tampilkan frekuensi bilangan x\n")
		fmt.Print("Pilih menu: ")
		fmt.Scanln(&menu)
		switch menu {
		case 1:
			fmt.Printf("%v\n", arr)
			break
		case 2:
			ganjil(&arr)
			break
		case 3:
			genap(&arr)
			break
		case 4:
			kelipatan(&arr)
			break
		case 5:
			hapus(&arr)
			break
		case 6:
			rata2(&arr)
			break
		case 7:
			simpangan_baku(&arr)
			break
		case 8:
			frekuensi(&arr)
			break
		default:
			fmt.Println("Masukkan nomor yang sesuai dengan menu!")
			break
		}
	}
}