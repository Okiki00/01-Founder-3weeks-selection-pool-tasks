package piscine

import "github.com/01-edu/z01"

//	func () {
//	  n := 5
//
// if n < 0 {
// do something
// }// else if n == 0 {
//
// } else {
// }
// }
func IsNegative(nb int) {
	if nb < 0 {
		z01.PrintRune('T')
	} else if nb >= 0 {
		z01.PrintRune('F')
	}
	z01.PrintRune('\n')
}
