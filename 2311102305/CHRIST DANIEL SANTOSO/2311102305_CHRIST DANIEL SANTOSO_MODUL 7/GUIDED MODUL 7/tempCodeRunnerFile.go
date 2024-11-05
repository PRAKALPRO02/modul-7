package main

// NAMA : CHRIST DANIEL SANTOSO
// NIM : 2311102305

import "fmt"

func main() {

	mahasiswa := map[string]string{
		"20231001": "Andi",
		"20231002": "Budi",
		"20231003": "Cici",
	}

	mahasiswa["20231004"] = "Dedi"

	fmt.Println("Daftar Mahasiswa:")
	fmt.Println("NIM\t\tNama")
	fmt.Println("-------------------------")
	for nim, nama := range mahasiswa {
		fmt.Printf("%s\t%s\n", nim, nama)
	}

	nim := "20231002"
	fmt.Println("\nNama Mahasiswa dengan NIM", nim, "adalah", mahasiswa[nim])

	delete(mahasiswa, "20231003")

	fmt.Println("\nDaftar Mahasiswa setelah dihapus:")
	fmt.Println("NIM\t\tNama")
	fmt.Println("-------------------------")
	for nim, nama := range mahasiswa {
		fmt.Printf("%s\t%s\n", nim, nama)
	}
}
