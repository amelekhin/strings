package shiftand

func getMask(pattern string) map[int]int {
	mask := make(map[int]int)
	patternLen := len(pattern)

	for i := range pattern {
		mask[int(pattern[i])] |= 1 << uint32(patternLen-i-1)
	}

	return mask
}

// Find finds all occurrences of substring "pattern"
// in a string "text" using the Shift-And method.
func Find(text, pattern string) []int {
	mask := getMask(pattern)
	patternLen := len(pattern)
	entries := make([]int, 0)

	high := 1 << uint32(patternLen-1)
	m := 0

	for i := range text {
		m = (m>>1 | high) & mask[int(text[i])]

		if (m & 1) != 0 {
			entries = append(entries, i-patternLen+1)
		}
	}

	return entries
}
