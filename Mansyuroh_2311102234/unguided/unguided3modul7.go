package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Membuat scanner untuk membaca input pengguna
	scanner := bufio.NewScanner(os.Stdin)

	// Meminta nama-nama klub
	fmt.Print("Masukkan nama Klub A: ")
	scanner.Scan()
	klubA := scanner.Text()

	fmt.Print("Masukkan nama Klub B: ")
	scanner.Scan()
	klubB := scanner.Text()

	// Variabel untuk menyimpan hasil pemenang
	var pemenang []string

	// Loop untuk input skor pertandingan
	for pertandingan := 1; ; pertandingan++ {
		// Input skor Klub A
		fmt.Printf("Masukkan skor %s pada pertandingan %d: ", klubA, pertandingan)
		scanner.Scan()
		skorA, errA := strconv.Atoi(scanner.Text())

		// Input skor Klub B
		fmt.Printf("Masukkan skor %s pada pertandingan %d: ", klubB, pertandingan)
		scanner.Scan()
		skorB, errB := strconv.Atoi(scanner.Text())

		// Cek jika skor negatif atau input tidak valid
		if errA != nil || errB != nil || skorA < 0 || skorB < 0 {
			fmt.Println("Pertandingan selesai")
			break
		}

		// Menentukan hasil pertandingan
		if skorA > skorB {
			fmt.Printf("Hasil %d: %s menang\n", pertandingan, klubA)
			pemenang = append(pemenang, klubA)
		} else if skorB > skorA {
			fmt.Printf("Hasil %d: %s menang\n", pertandingan, klubB)
			pemenang = append(pemenang, klubB)
		} else {
			fmt.Printf("Hasil %d: Draw\n", pertandingan)
			pemenang = append(pemenang, "Draw")
		}
	}

	// Menampilkan daftar hasil akhir pertandingan
	fmt.Println("\nRingkasan Hasil Pertandingan:")
	for i, hasil := range pemenang {
		fmt.Printf("Hasil %d: %s\n", i+1, hasil)
	}
}
