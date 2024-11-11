package main

import (
	"fmt"
)

func main() {
	var timA, timB string
	var scoreA, scoreB int
	var winners []string

	// Input club names
	fmt.Print("Tim A: ")
	fmt.Scanln(&timA)
	fmt.Print("Tim B: ")
	fmt.Scanln(&timB)

	matchNumber := 1

	for {
		fmt.Printf("Pertandingan %d : ", matchNumber)
		fmt.Scan(&scoreA, &scoreB)

		if scoreA < 0 || scoreB < 0 {
			fmt.Println("ERROR")
			break
		}

		if scoreA > scoreB {
			winners = append(winners, timA)
		} else if scoreB > scoreA {
			winners = append(winners, timB)
		} else {
			winners = append(winners, "Draw")
		}

		matchNumber++
	}

	fmt.Println("\nHasil Pertandingan:")
	for i, winner := range winners {
		fmt.Printf("Hasil %d: %s\n", i+1, winner)
	}

	fmt.Println("Program Selesai")
}
