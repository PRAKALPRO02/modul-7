package main

import (
	"fmt"
	"math"
)

type arrData [100]int

func InputData(Data *arrData, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("Masukan Data ke-%v =", i+1)
		fmt.Scan(&Data[i])
	}
}

func Display(Data *arrData, n int) {
	fmt.Print("{")
	for i := 0; i < n; i++ {
		fmt.Printf("%v ", Data[i])
	}
	fmt.Println("}")
}

func Genap(Data *arrData, n int) {
	fmt.Println("data indeks genap : ")
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Printf("indek ke- %v = %v\n", i, Data[i])
		}
	}
	fmt.Println()
}
func Ganjil(Data *arrData, n int) {
	fmt.Println("data Indeks ganjil: ")
	for i := 0; i < n; i++ {
		if i%2 != 0 {
			fmt.Printf("indek ke- %v = %v\n", i, Data[i])
		}
	}
	fmt.Println()
}
func Rata(Data *arrData, n int) float64 {
	sum := 0
	for i := 0; i < n; i++ {
		sum += Data[i]
	}
	mean := float64(sum / n)
	return mean
}

func simpanganBaku(Data *arrData, n int) {
	mean := Rata(Data, n)
	var variansisum float64
	variansisum = 0.0
	for i := 0; i < n; i++ {
		variansisum += math.Pow((float64(Data[i]) - mean), 2)
	}
	var Simpang float64
	Simpang = math.Sqrt(variansisum / float64(n-1))

	fmt.Printf("Simpangan Baku : %.2f\n", Simpang)
}

func CariIndeks(Data *arrData, n int) {
	var x int
	fmt.Print("Masukan indeks yang dicari : ")
	fmt.Scan(&x)
	if x > n {
		fmt.Println("Indeks yang kamu cari tidak ada")
	} else {
		fmt.Printf("Indeks %v dan kelipatannya \n", x)
		for i := 0; i < n; i++ {
			if i%x == 0 && i >= x {
				fmt.Printf("Indeks ke-%v = %v\n", i, Data[i])
			}
		}
	}
	fmt.Println()
}

func Hapus(Data *arrData, n int) {
	var index int
	fmt.Print("Masukan index array yang ingin dihapus : ")
	fmt.Scan(&index)
	if index >= 0 && index < n {
		Data := append(Data[:index], Data[index+1:]...)
		n--
		fmt.Print("Array setelah dihapus : ")
		for i := 0; i < n; i++ {
			fmt.Printf("%v ", Data[i])
		}
		fmt.Println()
	} else {
		fmt.Println("index tidak valid")
	}
}
func frekuensi(Data *arrData, n int) {
	F := 0
	var index int
	fmt.Print("Masukan index array yang ingin dihapus : ")
	fmt.Scan(&index)
	for i := 0; i < n; i++ {
		if Data[i] == index {
			F++
		}
	}
	fmt.Printf("Frekuensi dari %v adalah %v", index, F)

}

func main() {
	var Data arrData
	var n int
	pil := 1

	fmt.Print("Masukan Banyak Data: ")
	fmt.Scan(&n)
	InputData(&Data, n)
	for pil != 0 {

		fmt.Println("1. Tampilkan")
		fmt.Println("2. Tampilkan indeks ganjil")
		fmt.Println("3. Tampilkan indeks genap")
		fmt.Println("4. Tampilkan array dengan indeks kelipatan x")
		fmt.Println("5. Hapus")
		fmt.Println("6. Rata Rata")
		fmt.Println("7. Simpangan Baku")
		fmt.Println("8. Frekuensi")
		fmt.Println("Masukan Pilihan menu :")
		fmt.Scan(&pil)

		switch pil {
		case 1:
			Display(&Data, n)
			fmt.Println()
			break
		case 2:
			Ganjil(&Data, n)
			fmt.Println()
			break
		case 3:
			Genap(&Data, n)
			fmt.Println()
			break
		case 4:
			CariIndeks(&Data, n)
			fmt.Println()
			break
		case 5:
			Hapus(&Data, n)
			fmt.Println()
			break
		case 6:
			fmt.Println(Rata(&Data, n))
			fmt.Println()

			break
		case 7:
			simpanganBaku(&Data, n)
			fmt.Println()

			break
		case 8:
			frekuensi(&Data, n)
			fmt.Println()
			break
		default:
			fmt.Println("tidak ada pilihan")
		}
	}
}
