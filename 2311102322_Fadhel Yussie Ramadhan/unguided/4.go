package main 
 
import "fmt" 
 
const NMAX int = 127 
 
type tabel [NMAX]rune 
 
func isiArray(t *tabel, n *int) { 
    var input string
     fmt.Scanln(&input) 
    *n = 0 
    for _, ch := range input {
		         if ch == '.' || *n >= NMAX {
	             break 
        } 
        t[*n] = ch 
        *n++ 
    } } 
func cetakArray(t tabel, n int) { 
    for i := 0; i < n; i++ {
         fmt.Printf("%c", t[i]) 
    } 
    fmt.Println() 
} 
func balikkanArray(t *tabel, n int) { 
    for i := 0; i < n/2; i++ {
         t[i], t[n-i-1] = t[n-i-1], t[i] 
    } } func main() {
		
     var tab tabel
     var n int    
	 fmt.Print("Teks: ")     
	 isiArray(&tab, &n)     
	 balikkanArray(&tab, n)     
	 fmt.Print("Reverse Teks: ")     
	 cetakArray(tab, n) 
} 
