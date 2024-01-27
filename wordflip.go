package piscine

func WordFlip(str string) string {
	if str == "" {
		return "Invalid Output\n"
	}
	var s1 []string
	var s2 string
	var s string
	for _, v := range str {
		if v != ' ' {
			s2 += string(v)
		} else {
			s1 = append(s1, s2)
			s2 = ""
		}
	}
	if s2 != "" {
		s1 = append(s1, s2)
	}
	for i, j := 0, len(s1)-1; i < j; i, j = i+1, j-1 {
		s1[i], s1[j] = s1[j], s1[i]
	}
	for i := 0; i < len(s1); i++ {
		s += s1[i]
		if i != len(s1)-1 && len(s1[i+1]) != 0 {
			s += " "
		}
	}
	return s + "\n"
}
