package main

import (
	"github.com/01-edu/go-tests/lib/base"
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func AtoiBase(s string, base string) int {
	result := 0

	if isValidBase(base) {
		for _, sh := range s {
			index := 0
			for i, ch := range base {
				if sh == ch {
					index = i
					break
				}
			}
			result = result*len(base) + index
		}
		return result
	}

	return result
}

func isValidBase(base string) bool {

	charSet := make(map[rune]bool)
	for _, char := range base {
		if char == '+' || char == '-' {
			return false
		}
		if charSet[char] {
			return false
		}
		charSet[char] = true
	}

	return true
}
func main() {
	type node struct {
		s    string
		base string
	}

	args := []node{
		{s: "125", base: "0123456789"},
		{s: "1111101", base: "01"},
		{s: "7D", base: "0123456789ABCDEF"},
		{s: "uoi", base: "choumi"},
		{s: "bbbbbab", base: "-ab"},
	}

	for i := 0; i < 15; i++ {
		validBaseToInput := base.Valid()
		val := node{
			s:    base.StringFrom(validBaseToInput),
			base: validBaseToInput,
		}
		args = append(args, val)
	}
	for i := 0; i < 15; i++ {
		invalidBaseToInput := base.Invalid()
		val := node{
			s:    "thisinputshouldnotmatter",
			base: invalidBaseToInput,
		}
		args = append(args, val)
	}
	for _, arg := range args {
		challenge.Function("AtoiBase", AtoiBase, solutions.AtoiBase, arg.s, arg.base)
	}
}
