package main

import (
	"fmt"
)

// Fungsi untuk meminta input nama klub
func inputNamaKlub() (string, string) {
	var klubA, klubB string
	fmt.Print("Klub A: ")
	fmt.Scanln(&klubA)
	fmt.Print("Klub B: ")
	fmt.Scanln(&klubB)
	return klubA, klubB
}

// Fungsi untuk meminta input skor dari pertandingan dan mengembalikan skor serta status validitas
func inputSkor(pertandingan int, klubA, klubB string) (int, int, bool) {
	var skorA, skorB int
	fmt.Printf("Pertandingan %d: ", pertandingan)
	_, err := fmt.Scanln(&skorA, &skorB) // Menggunakan Scanln untuk input agar lebih jelas
	if err != nil {
		fmt.Println("Input tidak valid.")
		return 0, 0, false // Mengembalikan hasil gagal jika input tidak valid
	}

	if skorA < 0 || skorB < 0 {
		return skorA, skorB, false // skor tidak valid
	}
	fmt.Printf("// %s %d sedangkan %s %d\n", klubA, skorA, klubB, skorB)
	return skorA, skorB, true // skor valid
}

// Fungsi untuk menentukan pemenang berdasarkan skor
func tentukanPemenang(skorA, skorB int, klubA, klubB string) string {
	if skorA > skorB {
		return klubA
	} else if skorB > skorA {
		return klubB
	}
	return "Draw"
}

// Fungsi untuk menampilkan daftar hasil pertandingan
func tampilkanHasil(pemenang []string) {
	fmt.Println("Hasil pertandingan:")
	for i, hasil := range pemenang {
		fmt.Printf("Hasil %d: %s\n", i+1, hasil)
	}
	fmt.Println("Pertandingan selesai")
}

func main() {
	// Meminta input nama klub
	klubA, klubB := inputNamaKlub()

	// Array untuk menyimpan hasil pertandingan
	var pemenang []string
	var pertandingan int = 1

	for {
		// Meminta input skor
		skorA, skorB, valid := inputSkor(pertandingan, klubA, klubB)

		// Memeriksa validitas skor
		if !valid {
			fmt.Println("Skor tidak valid, program berhenti.")
			break
		}

		// Menentukan pemenang atau seri
		hasil := tentukanPemenang(skorA, skorB, klubA, klubB)
		pemenang = append(pemenang, hasil)
		pertandingan++

		// Memeriksa apakah perlu berhenti (misalnya jika sudah 12 pertandingan)
		if pertandingan > 12 {
			break
		}
	}

	// Menampilkan hasil pertandingan
	tampilkanHasil(pemenang)
}
