package shiftand

func getMask(pattern string) map[byte]byte {
	mask := make(map[byte]byte)
	patternLen := len(pattern)

	for i := range pattern {
		mask[pattern[i]] |= (1 << byte(patternLen-i-1))
	}

	return mask
}

// Find finds all occurrences of substring "pattern"
// in a string "text" using the Shift-And method.
func Find(text string, pattern string) []int {
	mask := getMask(pattern)
	patternLen := len(pattern)
	entries := make([]int, 0)
	high := byte(1 << byte(patternLen-1))
	m := byte(0)

	for i := range text {
		m = (m>>1 | high) & mask[text[i]]

		if (m & 1) != 0 {
			entries = append(entries, i-patternLen+1)
		}
	}

	return entries
}
