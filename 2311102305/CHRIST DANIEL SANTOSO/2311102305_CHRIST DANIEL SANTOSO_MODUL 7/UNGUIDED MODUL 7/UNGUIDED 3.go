package main

// NAMA : CHRIST DANIEL SANTOSO
// NIM : 2311102305

import "fmt"

func main() {
	var team1, team2 string
	var score1, score2 int
	var results []string

	fmt.Print("Nama Klub 1: ")
	fmt.Scanln(&team1)
	fmt.Print("Nama Klub 2: ")
	fmt.Scanln(&team2)

	gameCount := 1
	for {
		fmt.Printf("Skor Pertandingan %d (Format: Skor1 Skor2): ", gameCount)
		fmt.Scanln(&score1, &score2)

		if score1 < 0 || score2 < 0 {
			break
		}

		switch {
		case score1 > score2:
			results = append(results, team1)
		case score1 < score2:
			results = append(results, team2)
		default:
			results = append(results, "Draw")
		}

		gameCount++
	}

	fmt.Println("Semua pertandingan telah selesai")
	for i, result := range results {
		fmt.Printf("Hasil pertandingan %d: %s\n", i+1, result)
	}
}
