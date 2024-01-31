package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
)

func WeAreUnique(str1, str2 string) int {
	if str1 == "" && str2 == "" {
		return -1
	}
	s1 := make(map[rune]bool)
	s2 := make(map[rune]bool)
	for _, v := range str1 {
		s1[v] = true
	}
	for _, v := range str2 {
		s2[v] = true
	}
	i := 0
	for v := range s1 {
		if !s2[v] {
			i++
		}
	}
	for v := range s2 {
		if !s1[v] {
			i++
		}
	}
	return i
}

func main() {
	table := [][]string{
		{"abc", "def"},
		{"hello", "yoall"},
		{"everyone", ""},
		{"hello world", "fam"},
		{"abc", "abc"},
		{"", ""},
		{"pomme", "pomme"},
		{"+265", "265"},
		{"123231", "123231"},
		{"w^p@@j", "w^p@@j"},
		{"26235e5", "4478q92"},
		{"		", "		 "},
		{"AB$%d.52", "eepqdl.52"},
		{"", "eveRyone"},
		{"_55w1se", "55w1se"},
	}
	for _, arg := range table {
		challenge.Function("WeAreUnique", WeAreUnique, WeAreUnique, arg[0], arg[1])
	}
}
