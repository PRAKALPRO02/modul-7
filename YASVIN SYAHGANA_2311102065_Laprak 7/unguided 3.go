package main

import "fmt"

func main() {
	var clubA, clubB string
	var scoreA, scoreB int
	var results []string

	fmt.Print("Klub A: ")
	fmt.Scanln(&clubA)
	fmt.Print("Klub B: ")
	fmt.Scanln(&clubB)

	matchCount := 1

	for {
		fmt.Printf("Pertandingan %d: ", matchCount)
		fmt.Scan(&scoreA, &scoreB)

		if scoreA < 0 || scoreB < 0 {
			break
		}

		if scoreA > scoreB {
			results = append(results, fmt.Sprintf("Hasil %d: %s", matchCount, clubA))
		} else if scoreB > scoreA {
			results = append(results, fmt.Sprintf("Hasil %d: %s", matchCount, clubB))
		} else {
			results = append(results, fmt.Sprintf("Hasil %d: Draw", matchCount))
		}
		matchCount++
	}

	for _, result := range results {
		fmt.Println(result)
	}
	fmt.Println("Pertandingan selesai")
}
