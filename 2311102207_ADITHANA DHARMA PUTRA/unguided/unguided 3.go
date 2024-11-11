package main
import "fmt"

type Klub207 struct {
    Nama string
    Skor int
}

func main() {
    var Klub2071 Klub207
	var Klub2072 Klub207
    var pemenang []string
	var match int = 1

    fmt.Print("Masukkan nama Klub207 A: ")
    fmt.Scanln(&Klub2071.Nama)
	fmt.Print("Masukkan nama Klub207 B: ")
    fmt.Scanln(&Klub2072.Nama)

    for {
        fmt.Printf("pertandingan %v: ",match)
        fmt.Scanln(&Klub2071.Skor,&Klub2072.Skor)
		match++

        if Klub2071.Skor < 0 || Klub2072.Skor < 0 {
            break
        }

        if Klub2071.Skor > Klub2072.Skor {
            pemenang = append(pemenang, Klub2071.Nama)
        } else if Klub2072.Skor == Klub2071.Skor {
            pemenang = append(pemenang, "DRAW")
        } else if Klub2072.Skor > Klub2071.Skor {
            pemenang = append(pemenang, Klub2072.Nama)
        }
    }

    fmt.Println("Daftar Klub yang memenangkan pertandingan:")
    for _, Klub207 := range pemenang {
        fmt.Println(Klub207)
    }
}