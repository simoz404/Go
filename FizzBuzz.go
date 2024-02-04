func IntToString(n int) string {
	var s []rune
	for n > 0 {
		s = append([]rune{rune(n%10 + '0')}, s...)
		n /= 10
	}
	return string(s)
}

func fizzBuzz(n int) []string {
	var s []string
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			s = append(s, "FizzBuzz")
		} else if i%3 == 0 {
			s = append(s, "Fizz")
		} else if i%5 == 0 {
			s = append(s, "Buzz")
		} else {
			s = append(s, IntToString(i))
		}
	}
	return s
}