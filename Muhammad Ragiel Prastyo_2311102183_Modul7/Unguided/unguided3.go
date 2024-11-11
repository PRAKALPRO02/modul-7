// MUHAMMAD RAGIEL PRASTYO
// 2311102183
package main
import (
	"fmt"
)

func main() {
	var clubA, clubB string
	var scoreA, scoreB int

	fmt.Print("Klub A: ")
	fmt.Scanln(&clubA)
	fmt.Print("Klub B: ")
	fmt.Scanln(&clubB)

	var winners []string

	for i := 1; ; i++ {
		fmt.Printf("Pertandingan %d: ", i)
		fmt.Scanln(&scoreA, &scoreB)

		if scoreA < 0 || scoreB < 0 {
			break
		}

		if scoreA > scoreB {
			winners = append(winners, clubA)
		} else if scoreB > scoreA {
			winners = append(winners, clubB)
		} else {
			winners = append(winners, "Draw")
		}
	}

	fmt.Println("Hasil pertandingan:")
	for i, winner := range winners {
		fmt.Printf("Hasil %d: %s\n", i+1, winner)
	}
	fmt.Println("Pertandingan selesai")
}