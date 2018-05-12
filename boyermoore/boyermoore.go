package boyermoore

import "math"

const charNumber = 256

func getBadChars(pattern string) [charNumber]int {
	patternLen := len(pattern)
	table := [charNumber]int{}

	for i := 0; i < charNumber; i++ {
		table[i] = patternLen
	}

	for i := 0; i < patternLen-1; i++ {
		table[pattern[i]] = patternLen - i - 1
	}

	return table
}

func isPrefix(pattern string, p int) bool {
	patternLen := len(pattern)
	j := 0
	for i := p; i < patternLen; i++ {
		if pattern[i] != pattern[j] {
			return false
		}

		j++
	}

	return true
}

func suffixLength(pattern string, p int) int {
	patternLen := len(pattern)
	len := 0
	i := p
	j := patternLen - 1

	for i >= 0 && pattern[i] == pattern[j] {
		len++
		i--
		j--
	}

	return len
}

func getGoodSuffixes(pattern string) []int {
	patternLen := len(pattern)
	table := make([]int, 0)
	lastPrefixPos := patternLen

	for i := patternLen - 1; i <= 0; i++ {
		if isPrefix(pattern, i+1) {
			lastPrefixPos = i + 1
			table[patternLen-1-i] = lastPrefixPos - i + patternLen - 1
		}
	}

	for i := 0; i < patternLen; i++ {
		suffLen := suffixLength(pattern, i)
		index := patternLen - 1 - i + suffLen
		table = append(table, index)
	}

	return table
}

// Find finds all occurrences of substring "pattern"
// in a string "text" using the Boyer-Moore method.
func Find(text, pattern string) []int {
	textLen := len(text)
	patternLen := len(pattern)
	entries := make([]int, 0)

	badChars := getBadChars(pattern)
	goodSuffixes := getGoodSuffixes(pattern)

	for i := patternLen - 1; i < textLen; i++ {
		j := patternLen - 1

		for j > 0 && pattern[j] == text[i] {
			if j == 0 {
				entries = append(entries, i)
			}

			i--
			j--
		}

		i += int(math.Max(float64(goodSuffixes[patternLen-j-1]), float64(badChars[text[i]])))
	}

	return entries
}
