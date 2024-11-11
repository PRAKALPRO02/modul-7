package main

import "fmt"

func main() {
	// Membuat map dengan NIM sebagai kunci dan Nama sebagai nilai
	mahasiswa := map[string]string{
		"20231001": "Andi",
		"20231002": "Budi",
		"20231003": "Cici",
	}

	// Menambahkan data baru ke map
	mahasiswa["20231004"] = "Dedi"

	// Menampilkan seluruh isi map dalam format kolom dan baris
	fmt.Println("Daftar Mahasiswa:")
	fmt.Println("NIM\t\tNama")
	fmt.Println("-------------------------")
	for nim, nama := range mahasiswa {
		fmt.Printf("%s\t%s\n", nim, nama)
	}

	// Mengakses data berdasarkan NIM
	nim := "20231002"
	fmt.Println("\nNama Mahasiswa dengan NIM", nim, "adalah", mahasiswa[nim])

	// Menghapus data berdasarkan NIM
	delete(mahasiswa, "20231003")

	// Menampilkan isi map setelah data dihapus dalam format kolom dan baris
	fmt.Println("\nDaftar Mahasiswa setelah dihapus:")
	fmt.Println("NIM\t\tNama")
	fmt.Println("-------------------------")
	for nim, nama := range mahasiswa {
		fmt.Printf("%s\t%s\n", nim, nama)
	}
}