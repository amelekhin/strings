package border

func getLongestBorder(str string) int {
	strLen := len(str)
	borderIndex := 0

	for i := range str {
		prefix := str[0:i]
		suffix := str[strLen-i : strLen]

		if prefix == suffix {
			borderIndex = i
		}
	}

	return borderIndex
}

func getBorders(str string) []int {
	borders := make([]int, len(str))

	for i := range str {
		borders[i] = getLongestBorder(str[0 : i+1])
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
