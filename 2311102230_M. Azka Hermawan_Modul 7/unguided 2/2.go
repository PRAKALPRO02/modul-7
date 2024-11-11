package main

import "fmt"

func sqrt(x float64) float64 {
	z := x
	for i := 0; i < 1000; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	var N int
	fmt.Print("Masukkan jumlah elemen dalam array: ")
	fmt.Scan(&N)

	arr := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Printf("Masukkan elemen ke-%d: ", i)
		fmt.Scan(&arr[i])
	}

	for {
		fmt.Println("\nPilih list yang ingin anda cari: ")
		fmt.Println("1. Tampilkan seluruh elemen array")
		fmt.Println("2. Tampilkan elemen dengan indeks ganjil")
		fmt.Println("3. Tampilkan elemen dengan indeks genap")
		fmt.Println("4. Tampilkan elemen dengan indeks kelipatan bilangan x")
		fmt.Println("5. Hapus elemen pada indeks tertentu")
		fmt.Println("6. Tampilkan rata-rata bilangan dalam array")
		fmt.Println("7. Tampilkan simpangan baku bilangan dalam array")
		fmt.Println("8. Tampilkan frekuensi dari bilangan tertentu")
		fmt.Println("9. Keluar")

		var pil230 string
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pil230)

		switch pil230 {
		case "1":
			fmt.Println("Seluruh elemen array:", arr)

		case "2":
			fmt.Print("Elemen dengan indeks ganjil: ")
			for i := 1; i < N; i += 2 {
				fmt.Print(arr[i], " ")
			}
			fmt.Println()

		case "3":
			fmt.Print("Elemen dengan indeks genap: ")
			for i := 0; i < N; i += 2 {
				fmt.Print(arr[i], " ")
			}
			fmt.Println()

		case "4":
			var x int
			fmt.Print("Masukkan bilangan x: ")
			fmt.Scan(&x)
			fmt.Printf("Elemen dengan indeks kelipatan %d: ", x)
			for i := 0; i < N; i++ {
				if i%x == 0 {
					fmt.Print(arr[i], " ")
				}
			}
			fmt.Println()

		case "5":
			var idx int
			fmt.Print("Masukkan indeks yang ingin dihapus: ")
			fmt.Scan(&idx)
			if idx >= 0 && idx < N {
				arr = append(arr[:idx], arr[idx+1:]...)
				N--
				fmt.Println("Array setelah penghapusan:", arr)
			} else {
				fmt.Println("Indeks tidak valid!")
			}

		case "6":
			total := 0
			for _, val := range arr {
				total += val
			}
			rataRata := float64(total) / float64(N)
			fmt.Println("Rata-rata nilai array:", rataRata)

		case "7":
			total := 0
			for _, val := range arr {
				total += val
			}
			rataRata := float64(total) / float64(N)

			var sumSquare float64
			for _, val := range arr {
				sumSquare += (float64(val) - rataRata) * (float64(val) - rataRata)
			}
			simpanganBaku := sqrt(sumSquare / float64(N))
			fmt.Println("Simpangan baku dari array:", simpanganBaku)

		case "8":
			var target int
			fmt.Print("Masukkan bilangan yang ingin dihitung frekuensinya: ")
			fmt.Scan(&target)
			frekuensi := 0
			for _, val := range arr {
				if val == target {
					frekuensi++
				}
			}
			fmt.Printf("Frekuensi bilangan %d dalam array: %d\n", target, frekuensi)

		case "9":
			fmt.Println("Keluar.")
			return

		default:
			fmt.Println("pilihan yang anda pilih tidak valid.")
		}
	}
}
