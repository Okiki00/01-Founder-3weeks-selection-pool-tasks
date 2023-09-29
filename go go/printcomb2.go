package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := 0; i <= 98; i++ {
		for j := i + 1; j <= 99; j++ {
			printTwoDigits(i)
			z01.PrintRune(' ')
			printTwoDigits(j)
			if i != 98 || j != 99 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}

func printTwoDigits(num int) {
	rune1 := '0' + rune(num/10)
	rune2 := '0' + rune(num%10)
	z01.PrintRune(rune1)
	z01.PrintRune(rune2)
}
