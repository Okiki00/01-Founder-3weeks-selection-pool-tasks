package main

import "github.com/01-edu/z01"

func main() {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	// for i := 'a'; i <= 'z'; i++ {
	// 	z01.PrintRune(i)
	// }
	for i := 0; i < len(alphabet); i++ {
		z01.PrintRune(rune(alphabet[i]))
	}
	z01.PrintRune('\n')
}
