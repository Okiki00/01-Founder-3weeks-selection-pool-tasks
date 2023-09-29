package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
	z int
	a int
}

func setPoint(ptr *point) {
	ptr.x = 52
	ptr.y = 50
	ptr.z = 50
	ptr.a = 49
}

func main() {
	points := &point{}
	setPoint(points)
	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	z01.PrintRune(rune(points.x))
	z01.PrintRune(rune(points.y))
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	z01.PrintRune(rune(points.z))
	z01.PrintRune(rune(points.a))
	z01.PrintRune('\n')
}
