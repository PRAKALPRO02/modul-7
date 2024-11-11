package main

import "fmt"

func skorBola(arr []string) {
	var skor1, skor2 int
	var i int = 0
	fmt.Print("masukan klub 1: ")
	fmt.Scan(&arr[0])
	fmt.Print("masukan klub 2: ")
	fmt.Scan(&arr[1])

	fmt.Println()

	for {
		i++
		fmt.Print("pertandingan ", i, ": ")
		fmt.Scan(&skor1, &skor2)

		if skor1 < 0 || skor2 < 0 {
			break
		}

		if skor1 > skor2 {
			fmt.Println("hasil ", i, ": ", arr[0])
		} else if skor2 > skor1 {
			fmt.Println("hasil ", i, ": ", arr[1])
		} else if skor1 == skor2 || skor2 == skor1 {
			fmt.Println("hasil ", i, ": Draw")
		}

	}
}

func main() {
	arr := make([]string, 2)
	skorBola(arr)
}