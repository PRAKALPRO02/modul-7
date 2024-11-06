package main


import (

	"fmt"
	"math"
)

func jarak(x1, y1, x2, y2 int) float64 {
	return math.Sqrt(float64((x2-x1)*(y2-1) + (y2-y1)*(y2-y1)))
}

func didalamLingkaran(x, y, cx, cy, r int) bool{
	distance := jarak(x, y, cx, cy)
	return distance <= float64(r)
}

func main(){
	var cx1, cy1, r1 int
	fmt.Println("Masukkan koordinat x, y, dan radius dalam lingkaran 1 : ")
	fmt.Scan(&cx1, &cy1, &r1)

	var cx2, cy2, r2 int
	fmt.Println("Masukkan koordinat x, y, dan radius untuk lingkaran 2 : ")
	fmt.Scan(&cx2, &cy2, &r2)

	var x, y int
	fmt.Println("Masukkan koordinat x dan y untuk titik sembarang : ")
	fmt.Scan(&x, &y)

	didalamLingkaran1 := didalamLingkaran(x, y, cx1, cy1, r1)
	didalamLingkaran2 := didalamLingkaran(x, y, cx2, cy2, r2)

	if didalamLingkaran1 && didalamLingkaran2 {
		fmt.Println("Titik didalalm lingkaran 1 dan 2")
	} else if didalamLingkaran1{
		fmt.Println("Titik didalam lingkaran 1")
	} else if didalamLingkaran2 {
		fmt.Println("Titik didalam lingkaran 2")

	} else {
		fmt.Println("Titik diluar lingkaran 1 dan 2")
	}
}