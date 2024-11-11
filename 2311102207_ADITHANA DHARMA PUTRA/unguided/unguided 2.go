package main

import "fmt"
import "math"

func tampilkanSeluruharray(array207 []int) {
	fmt.Println("Seluruh elemen array:", array207)
}

func tampilkanIndeksGanjil(array207 []int) {
	fmt.Print("Elemen dengan indeks ganjil: ")
	for i := 1; i < len(array207); i += 2 {
		fmt.Print(array207[i], " ")
	}
	fmt.Println()
}

func tampilkanIndeksGenap(array207 []int) {
	fmt.Print("Elemen dengan indeks genap: ")
	for i := 0; i < len(array207); i += 2 {
		fmt.Print(array207[i], " ")
	}
	fmt.Println()
}

func tampilkandengankelipatan(array207 []int, kelipatan int) {
	fmt.Printf("Elemen yang merupakan kelipatan dari %v :", kelipatan)
	for i := 0; i < len(array207); i += kelipatan {
		fmt.Print(array207[i], " ")
	}
	fmt.Println()
}

func tampilkanFrekuensi(array207 []int, x int) {
	frekuensi := 0
	for _, angka := range array207 {
		if angka == x {
			frekuensi++
		}
	}
	fmt.Printf("Frekuensi dari %d: %d\n", x, frekuensi)
}

func hapusElemen(array207 []int, x int) []int {
	result := []int{}
	for _, elemen := range array207 {
		if elemen != x {
			result = append(result, elemen)
		}
	}
	return result
}

func hitungRataRata(array207 []int) float64 {
	if len(array207) == 0 {
		return 0
	}
	total := 0
	for _, angka := range array207 {
		total += angka
	}
	return float64(total) / float64(len(array207))
}

func hitungSimpanganBaku(array207 []int) float64 {
	if len(array207) == 0 {
		return 0
	}
	rataRata := hitungRataRata(array207)
	var total float64
	for _, angka := range array207 {
		total += math.Pow(float64(angka)-rataRata, 2)
	}
	return math.Sqrt(total / float64(len(array207)))
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scan(&n)
	array207 := make([]int, n)
	fmt.Print("Masukkan array: ")
	for i := 0; i < n; i++ {
		fmt.Scan(&array207[i])
	}

	tampilkanSeluruharray(array207)
	tampilkanIndeksGanjil(array207)
	tampilkanIndeksGenap(array207)

	var kelipatan int
	fmt.Print("Masukkan kelipatan: ")
	fmt.Scan(&kelipatan)
	tampilkandengankelipatan(array207, kelipatan)

	tampilkanFrekuensi(array207, 3)

	rataRata := hitungRataRata(array207)
	fmt.Printf("Rata-rata dari array: %.2f\n", rataRata)
	
	simpanganBaku := hitungSimpanganBaku(array207)
	fmt.Printf("Simpangan baku dari array: %.2f\n", simpanganBaku)
	
	var elemenHapus int
	fmt.Print("Masukkan elemen yang ingin dihapus: ")
	fmt.Scan(&elemenHapus)

	array207 = hapusElemen(array207, elemenHapus)
	tampilkanSeluruharray(array207)
}