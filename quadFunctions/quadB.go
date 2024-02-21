package piscine

import (
	"github.com/01-edu/z01"
)

func QuadB(x, y int) {
	if x >= 0 && y >= 0 {
		for a := 0; a < x; a++ {
			if a == 0 {
				z01.PrintRune('/')
			} else if a == x-1 {
				z01.PrintRune('\\')
			} else {
				z01.PrintRune('*')
			}
		}
		z01.PrintRune('\n')

		for b := 0; b < y-2; b++ {
			z01.PrintRune('*')
			for a := 0; a < x-2; a++ {
				z01.PrintRune(' ')
			}
			if !(x == 1) {
				z01.PrintRune('*')
			}
			z01.PrintRune('\n')

		}

		if y > 1 {
			for a := 0; a < x; a++ {
				if a == 0 {
					z01.PrintRune('\\')
				} else if a == x-1 {
					z01.PrintRune('/')
				} else {
					z01.PrintRune('*')
				}
			}
			z01.PrintRune('\n')
		}
	}
}
