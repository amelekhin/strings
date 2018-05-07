package border

func getBorders(str string) []int {
	strLen := len(str)
	borders := make([]int, strLen)

	t := 0
	for i := range str {
		if i == 0 {
			continue
		}

		for t > 0 && str[i] != str[t] {
			t = borders[t-1]
		}

		if str[i] == str[t] {
			t++
		}

		borders[i] = t
	}

	return borders
}

// Find finds all occurrences of substring "pattern"
// in a string "text" using the "borders" method.
func Find(text string, pattern string) []int {
	borders := getBorders(pattern + "$" + text)
	entries := make([]int, 0)
	patternLen := len(pattern)

	for i, border := range borders {
		if border == patternLen {
			// patternLen is multiplied by 2 to remove the offset
			entries = append(entries, i-patternLen*2)
		}
	}

	return entries
}
