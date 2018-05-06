package block

func getBlocks(str string) []int {
	strLen := len(str)
	blocks := make([]int, strLen)
	l, r := 0, 0

	for i := range str {
		if i > r {
			r = i
			l = r

			for r < strLen && str[r-l] == str[r] {
				r++
			}

			blocks[i] = r - l
			r--
		} else {
			k := i - l
			if blocks[k] < r-i+1 {
				blocks[i] = blocks[k]
			} else {
				l = i
				for r < strLen && str[r-l] == str[r] {
					r++
				}

				blocks[i] = r - l
				r--
			}
		}
	}

	return blocks
}

// Find finds all occurrences of substring "pattern"
// in a string "text" using the "blocks" method.
func Find(text string, pattern string) []int {
	blocks := getBlocks(pattern + "$" + text)
	entries := make([]int, 0)
	patternLen := len(pattern)

	for i := range blocks {
		if blocks[i] == patternLen {
			entries = append(entries, i-patternLen-1)
		}
	}

	return entries
}
