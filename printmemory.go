package main

import (
	"github.com/01-edu/z01"
)

func PrintHex(n int) {
	var a []int
	for n > 0 {
		a = append(a, n%16)
		n /= 16
	}
	base16 := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}

	if len(a) == 1 {
		z01.PrintRune('0')
	}
	for i := len(a) - 1; i >= 0; i-- {
		z01.PrintRune(base16[a[i]])
	}
}

func PrintMemory(arr [10]byte) {
	for i, v := range arr {
		if v == 0 {
			z01.PrintRune('0')
			z01.PrintRune('0')
		}
		if i < 4 {
			PrintHex(int(v))
			if i != 3 {
				z01.PrintRune(' ')
			} else {
				z01.PrintRune('\n')
			}
		} else if i < 8 {
			PrintHex(int(v))
			if i != 7 {
				z01.PrintRune(' ')
			} else {
				z01.PrintRune('\n')
			}
		} else if i < 10 {
			PrintHex(int(v))
			if i != 9 {
				z01.PrintRune(' ')
			} else {
				z01.PrintRune('\n')
			}
		}
	}
	for _, v := range arr {
		if v < 32 || v > 126 {
			z01.PrintRune('.')
		} else {
			z01.PrintRune(rune(v))
		}
	}
	z01.PrintRune('\n')
}

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}
