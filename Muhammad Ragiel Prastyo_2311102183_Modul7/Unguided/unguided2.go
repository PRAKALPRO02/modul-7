// MUHAMMAD RAGIEL PRASTYO
// 2311102183
package main
import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scan(&n)

	array := make([]int, n)
	fmt.Println("Masukkan elemen-elemen array:")
	for i := 0; i < n; i++ {
		fmt.Printf("Elemen %d: ", i)
		fmt.Scan(&array[i])
	}

	fmt.Println("a. Isi keseluruhan array:", array)

	fmt.Print("b. Elemen-elemen array dengan indeks ganjil: ")
	for i := 1; i < n; i += 2 {
		fmt.Print(array[i], " ")
	}
	fmt.Println()

	fmt.Print("c. Elemen-elemen array dengan indeks genap: ")
	for i := 0; i < n; i += 2 {
		fmt.Print(array[i], " ")
	}
	fmt.Println()

	var x int
	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&x)
	fmt.Printf("d. Elemen-elemen array dengan indeks kelipatan %d: ", x)
	for i := 0; i < n; i += x {
		fmt.Print(array[i], " ")
	}
	fmt.Println()

	var index int
	fmt.Print("Masukkan indeks yang ingin dihapus: ")
	fmt.Scan(&index)
	if index >= 0 && index < n {
		array = append(array[:index], array[index+1:]...)
		fmt.Println("e. Isi array setelah penghapusan:", array)
	} else {
		fmt.Println("Indeks tidak valid.")
	}

	sum := 0
	for _, value := range array {
		sum += value
	}
	average := float64(sum) / float64(len(array))
	fmt.Printf("f. Rata-rata dari bilangan dalam array: %.2f\n", average)

	varianceSum := 0.0
	for _, value := range array {
		varianceSum += math.Pow(float64(value)-average, 2)
	}
	variance := varianceSum / float64(len(array))
	stdDev := math.Sqrt(variance)
	fmt.Printf("g. Standar deviasi dari bilangan dalam array: %.2f\n", stdDev)

	var target int
	fmt.Print("Masukkan bilangan yang ingin dicari frekuensinya: ")
	fmt.Scan(&target)
	frequency := 0
	for _, value := range array {
		if value == target {
			frequency++
		}
	}
	fmt.Printf("h. Frekuensi dari bilangan %d dalam array: %d\n", target, frequency)
}