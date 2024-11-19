 // Fajar Farizqi Azmi
 // 2311102192 

 
package main

import (
	"fmt"
	"math"
)

// Struct untuk merepresentasikan titik
type Titik struct {
	x, y int
}

// Struct untuk merepresentasikan lingkaran
type Lingkaran struct {
	titikPusat Titik
	jariJari   int
}

// Fungsi untuk mengecek apakah suatu titik berada di dalam lingkaran
func apakahTitikDiDalamLingkaran(t Titik, l Lingkaran) bool {
	jarak := math.Sqrt(math.Pow(float64(t.x-l.titikPusat.x), 2) + math.Pow(float64(t.y-l.titikPusat.y), 2))
	return jarak <= float64(l.jariJari)
}

func main() {
	// Masukan baris pertama untuk lingkaran 1
	var x1, y1, r1 int
	fmt.Printf("Masukkan untuk lingkaran 1: ")
	fmt.Scan(&x1, &y1, &r1)
	lingkaran1 := Lingkaran{Titik{x1, y1}, r1}

	// Masukan baris kedua untuk lingkaran 2
	var x2, y2, r2 int
	fmt.Printf("Masukkan untuk lingkaran 2: ")
	fmt.Scan(&x2, &y2, &r2)
	lingkaran2 := Lingkaran{Titik{x2, y2}, r2}

	// Masukan baris ketiga untuk titik sembarang
	var x, y int
	fmt.Printf("Masukkan untuk titik sembarang: ")
	fmt.Scan(&x, &y)
	titik := Titik{x, y}

	// Mengecek apakah titik berada di dalam lingkaran 1, lingkaran 2, atau keduanya
	diLingkaran1 := apakahTitikDiDalamLingkaran(titik, lingkaran1)
	diLingkaran2 := apakahTitikDiDalamLingkaran(titik, lingkaran2)

	// Menentukan keluaran berdasarkan hasil pengecekan
	if diLingkaran1 && diLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if diLingkaran1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if diLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}