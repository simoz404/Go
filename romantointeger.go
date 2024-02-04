func RomanToInt(n rune) int {
	switch n {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	}
	return 0
}

func romanToInt(s string) int {
	result := 0
	s1 := []rune(s)

	for i := 0; i < len(s1); i++ {
		v := RomanToInt(s1[i])

		if i+1 < len(s1) && RomanToInt(s1[i]) < RomanToInt(s1[i+1]) {
			result -= v
		} else {
			result += v
		}
	}

	return result
}
