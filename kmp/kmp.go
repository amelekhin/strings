package kmp

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

func getModifiedPatternBorders(pattern string) []int {
	patternBorders := getBorders(pattern)
	patternBordersLen := len(patternBorders)
	modifiedPatternBorders := make([]int, patternBordersLen)

	for i := range patternBorders {
		if i < patternBordersLen-1 && pattern[patternBorders[i]] != pattern[i+1] {
			modifiedPatternBorders[i] = patternBorders[i]
		} else {
			modifiedPatternBorders[i] = modifiedPatternBorders[i-1]
		}
	}

	return modifiedPatternBorders
}

// Find finds all occurrences of substring "pattern"
// in a string "text" using the Knuth-Morris-Pratt method.
func Find(text string, pattern string) []int {
	modifiedPatternBorders := getModifiedPatternBorders(pattern)
	patternLen := len(pattern)
	entries := make([]int, 0)

	t := 0
	for i := range text {
		for t > 0 && t < patternLen && pattern[t] != text[i] {
			t = modifiedPatternBorders[t-1]
		}

		if t < patternLen && pattern[t] == text[i] {
			t++
		}

		if t == patternLen {
			entries = append(entries, i-patternLen+1)
			t = 0
		}
	}

	return entries
}

// FindRT finds all occurrences of substring "pattern"
// in a string "text" using the Knuth-Morris-Pratt (relatime) method.
func FindRT(text string, pattern string) []int {
	const radix = 256
	patternLen := len(pattern)
	patternBorders := make([][]int, radix)

	for i := 0; i < radix; i++ {
		patternBorders[i] = make([]int, patternLen)
	}

	patternBorders[pattern[0]][0] = 1

	entries := make([]int, 0)
	k := 0

	for i := 1; i < patternLen; i++ {
		for j := 0; j < radix; j++ {
			patternBorders[j][i] = patternBorders[j][k]
		}

		patternBorders[pattern[i]][i] = i + 1
		k = patternBorders[pattern[i]][k]
	}

	j := 0
	for i := 0; i < len(text); i++ {
		j = patternBorders[text[i]][j]

		if j == patternLen {
			entries = append(entries, i-patternLen+1)
			j = 0
		}
	}

	return entries
}
