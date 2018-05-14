package boyermoore

import "math"

const alphabetSize = 256

func getBadChars(pattern string) [alphabetSize - 1]int {
	var badChars [alphabetSize - 1]int
	patternLen := len(pattern)

	for i := 0; i < patternLen; i++ {
		badChars[pattern[i]] = i
	}

	return badChars
}

func getGoodSuffixes(pattern string) []int {
	patternLen := len(pattern)
	bPos := make([]int, patternLen+1, patternLen+1)
	goodSuffixes := make([]int, patternLen+1, patternLen+1)

	i := patternLen
	j := patternLen + 1
	bPos[i] = j

	for i > 0 {
		for j > 0 && j < patternLen+1 && pattern[i-1] != pattern[j-1] {
			if goodSuffixes[j] == 0 {
				goodSuffixes[j] = j - i
			}

			j = bPos[j]
		}

		i--
		j--
		bPos[i] = j
	}

	j = bPos[0]
	for i = 0; i < patternLen+1; i++ {
		if goodSuffixes[i] == 0 {
			goodSuffixes[i] = j
		}

		if i == j {
			j = bPos[j]
		}
	}

	return goodSuffixes
}

// FindBC finds all occurrences of substring "pattern"
// in a string "text" using the Boyer-Moore method (bad char shifts).
func FindBC(text, pattern string) []int {
	textLen := len(text)
	patternLen := len(pattern)
	badChars := getBadChars(pattern)
	entries := make([]int, 0)

	i := 0
	for i <= textLen-patternLen {
		j := patternLen - 1

		for j >= 0 && text[i+j] == pattern[j] {
			j--
		}

		if j < 0 {
			entries = append(entries, i)
		}

		shift := j - badChars[text[i+j]]
		i += int(math.Max(1, float64(shift)))
	}

	return entries
}

// FindGS finds all occurrences of substring "pattern"
// in a string "text" using the Boyer-Moore method (good suffixes shifts).
func FindGS(text, pattern string) []int {
	textLen := len(text)
	patternLen := len(pattern)
	goodSuffixes := getGoodSuffixes(pattern)
	entries := make([]int, 0)

	i := 0
	for i <= textLen-patternLen {
		j := patternLen - 1

		for j >= 0 && text[i+j] == pattern[j] {
			j--
		}

		if j < 0 {
			entries = append(entries, i)
			i += goodSuffixes[0]
		} else {
			shift := goodSuffixes[j+1]
			i += int(math.Max(1, float64(shift)))
		}
	}

	return entries
}
