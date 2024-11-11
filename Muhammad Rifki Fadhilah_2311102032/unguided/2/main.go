package main

import(
  "fmt"
  "math"
)

func remove(array []int, s int) []int{
  return append(array[:s], array[s+1:]...)
}


func average(arr []int, n int)float64{
  sum := 0
  for i := 0; i < n ; i++ {
    sum += (arr[i])
  }
  var avg float64 = float64(sum)/float64(n)
  return avg
}

func deviasi(arr []int) float64 {
	var sum, mean, sd float64
	n := float64(len(arr))

	// Konversi setiap elemen dari int ke float64 untuk perhitungan deviasi
	floatArr := make([]float64, len(arr))
	for i, value := range arr {
		floatArr[i] = float64(value)
		sum += floatArr[i]
	}

	mean = sum / n

	for _, value := range floatArr {
		sd += math.Pow(value-mean, 2)
	}

	sd = math.Sqrt(sd / n)
	return sd
}

func frekuensi(arr []int, bil int) int {
	count := 0
	for _, item := range arr {
		if item == bil {
			count++
		}
	}
	return count
}

func main(){
  var n int
  fmt.Print("Masukkan jumlah elemen array: ")
  fmt.Scan(&n)

  arr := make([]int, n)
  fmt.Println("Masukkan elemen-elemen array: ")
  for i := 0; i < n; i++ {
    fmt.Scan(&arr[i])
  }

  //Menampilkan keseluruhan isi array 
  fmt.Println("Keseluruhan elemen dalam array: ", arr)
  fmt.Println()

  //Menampilkan bilangan genap dalam array
  fmt.Print("Bilangan genap dalam array: ")
  for i := 0; i < n; i++ {
    if arr[i]%2 == 0 {
      fmt.Print(arr[i]," ")
    }
  }
  fmt.Println()

  //Menampilkan bilangan ganjil dalam array
  fmt.Print("Bilangan ganjil dalam array: ")
  for i := 0; i < n; i++ {
    if arr[i]%2 == 1 {
      fmt.Print(arr[i]," ")
    }
  }
  fmt.Println()

  //Menampilkan bilangan dengan indeks kelipatan X
  var x int
  fmt.Print("Masukkan nilai x untuk indeks kelipatan: ")
  fmt.Scan(&x)

  if x <= 0 && x >= n {
    fmt.Print("Nilai x tidak valid")
  }

  fmt.Print("Elemen pada indeks kelipatan ", x, " : ")
  for i := x; i < n; i+= x {
    fmt.Print(arr[i]," ")
  }
  fmt.Println()

  //Menghapus elemen array dengan indeks tertentu dan menampilkan keseluruhannya
  var index int
  fmt.Print("Masukkan posisi indeks yang ingin dihapus: ")
  fmt.Scan(&index)

  if index >= 0 && index < n{
    arr = remove(arr,index)
    fmt.Print("Isi array setelah dihapus: ", arr)
  }else{
    fmt.Print("Indeks tidak valid")
  }
  fmt.Println()

  fmt.Printf("Rata-rata bilangan di array: %2f\n", average(arr,len(arr)))
  
  //Menampilkan standar deviasi dari bilangan dalam array
  fmt.Printf("Standar deviasi bilangan di array: %2f\n", deviasi(arr))
  
  //Menampilkan frekuensi dari suatu bilangan dalam array
  var bil int
	fmt.Print("Masukkan bilangan yang ingin dicari frekuensinya: ")
	fmt.Scan(&bil)
	fmt.Printf("Frekuensi bilangan %d dalam array: %d\n", bil, frekuensi(arr, bil))
}