package main

// NAMA : CHRIST DANIEL SANTOSO
// NIM : 2311102305

import "fmt"

const maxSize int = 127

type ArrayChar struct {
	data [maxSize]rune
	size int
}

func inputArray(arr *ArrayChar, length *int) {
	fmt.Print("Masukkan karakter (akhiri dengan titik): ")
	var char rune
	for *length = 0; *length < maxSize; *length++ {
		fmt.Scanf("%c", &char)
		if char == '.' {
			break
		}
		arr.data[*length] = char
	}
	arr.size = *length
}

func printArray(arr ArrayChar, length int) {
	for i := 0; i < length; i++ {
		fmt.Printf("%c", arr.data[i])
	}
	fmt.Println()
}

func reverseArray(arr *ArrayChar, length int) {
	for i := 0; i < length/2; i++ {
		arr.data[i], arr.data[length-i-1] = arr.data[length-i-1], arr.data[i]
	}
}

func isPalindrome(arr ArrayChar, length int) bool {
	for i := 0; i < length/2; i++ {
		if arr.data[i] != arr.data[length-i-1] {
			return false
		}
	}
	return true
}

func main() {
	var arr ArrayChar
	var length int

	inputArray(&arr, &length)

	fmt.Print("Teks: ")
	printArray(arr, length)

	fmt.Print("Reverse teks: ")
	reverseArray(&arr, length)
	printArray(arr, length)

	if isPalindrome(arr, length) {
		fmt.Println("Palindrom: true")
	} else {
		fmt.Println("Palindrom: false")
	}
}
