package main

import "fmt"

func main() {
	var nama_a, nama_b string
	var a, b int
	var arr []string

	fmt.Print("Klub A : ")
	fmt.Scan(&nama_a)
	fmt.Print("Klub B : ")
	fmt.Scan(&nama_b)

	for i := 0; ; i++ {
		fmt.Print("Pertandingan ", i+1, ": ")
		fmt.Scan(&a, &b)
		if a < 0 || b < 0 {
			break
		}
		if a > b {
			arr = append(arr, nama_a)
		} else if a < b {
			arr = append(arr, nama_b)
		} else {
			arr = append(arr, "Draw")
		}
	}

	for i := 0; i < len(arr); i++ {
		fmt.Println("Hasil", i+1, ":", arr[i])
	}
}
