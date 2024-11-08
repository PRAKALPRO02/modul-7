package main

import (
        "fmt"
)

func main() {
        var klubA, klubB string
        var skorA, skorB int

        fmt.Print("Klub A: ")
        fmt.Scan(&klubA)
        fmt.Print("Klub B: ")
        fmt.Scan(&klubB)

        var pemenang []string

        for {
                fmt.Printf("Pertandingan: ")
                fmt.Scan(&skorA, &skorB)

                if skorA < 0 || skorB < 0 {
                        break
                }

                if skorA > skorB {
                        pemenang = append(pemenang, klubA)
                } else if skorB > skorA {
                        pemenang = append(pemenang, klubB)
                } else {
                        pemenang = append(pemenang, "Draw")
                }
        }

        fmt.Println("Hasil pertandingan:")
        for i, p := range pemenang {
                fmt.Printf("Hasil %d: %s\n", i+1, p)
        }
}