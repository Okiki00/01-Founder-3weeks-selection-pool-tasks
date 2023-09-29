package main

import "github.com/01-edu/z01"

func main() {
	number := "0123456789"
	for i := 0; i < len(number); i++ {
		z01.PrintRune(rune(number[i]))
	}
	z01.PrintRune('\n')
}
