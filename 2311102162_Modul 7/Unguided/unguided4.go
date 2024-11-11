// 2311102162
package main

import (
	"fmt"
	"strings"
)

func reverse(text string) string {
	runes := []rune(text)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func isPalindrome(text string) bool {
	reversedText := reverse(text)
	return strings.ToLower(text) == strings.ToLower(reversedText)
}

func main() {
	var text string

	fmt.Print("Masukkan teks: ")
	fmt.Scanln(&text)

	reversedText := reverse(text)
	fmt.Println("Reverse teks:", reversedText)

	if isPalindrome(text) {
		fmt.Println("Palindrom: true")
	} else {
		fmt.Println("Palindrom: false")
	}
}
