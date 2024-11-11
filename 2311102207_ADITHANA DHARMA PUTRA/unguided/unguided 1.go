package main

import "fmt"
import "math"

type data2311102207 struct {
	titikX float64
	titikY float64
	cx     float64
	cy     float64
	r      float64
}

func jarak(titikX, titikY float64, lingkaran1, lingkaran2 data2311102207) string {
	jarakLingkaran1 := math.Sqrt(math.Pow(titikX-lingkaran1.cx, 2) + math.Pow(titikY-lingkaran1.cy, 2))
	jarakLingkaran2 := math.Sqrt(math.Pow(titikX-lingkaran2.cx, 2) + math.Pow(titikY-lingkaran2.cy, 2))

	if jarakLingkaran1 <= lingkaran1.r && jarakLingkaran2 <= lingkaran2.r{
		return "titik didalam lingkaran 1 dan 2"
	} else if jarakLingkaran1 <= lingkaran1.r {
		return "titik didalam lingkaran 1"
	} else if jarakLingkaran2 <= lingkaran2.r {
		return "titik didalam lingkaran 2"
	} else {
		return "titik diluar lingkaran 1 dan 2"
	}
}

func main() {
	var lingkaran1, lingkaran2 data2311102207

	fmt.Print("Masukkan pusat koordinat dan radius lingkaran 1 (x y radius): ")
	fmt.Scan(&lingkaran1.cx, &lingkaran1.cy, &lingkaran1.r)

	fmt.Print("Masukkan pusat koordinat dan radius lingkaran 2 (x y radius): ")
	fmt.Scan(&lingkaran2.cx, &lingkaran2.cy, &lingkaran2.r)

	fmt.Print("Masukkan titik koordinat (x y): ")
	fmt.Scan(&lingkaran1.titikX, &lingkaran1.titikY)

	hasil := jarak(lingkaran1.titikX, lingkaran1.titikY, lingkaran1, lingkaran2)
	fmt.Println(hasil)
}