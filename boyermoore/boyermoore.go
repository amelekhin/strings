package boyermoore

import "math"

const alphabetSize = 256

func getBadCharShifts(pattern string) [alphabetSize - 1]int {
	var badChars [alphabetSize - 1]int
	patternLen := len(pattern)

	for i := 0; i < patternLen; i++ {
		badChars[pattern[i]] = i
	}

	return badChars
}

func getGoodSuffixShifts(pattern string) []int {
	patternLen := len(pattern)
	borderPos := make([]int, patternLen+1, patternLen+1)
	goodSuffixes := make([]int, patternLen+1, patternLen+1)

	i := patternLen
	j := patternLen + 1
	borderPos[i] = j

	// Step 1:
	for i > 0 {
		for j <= patternLen && pattern[i-1] != pattern[j-1] {
			if goodSuffixes[j] == 0 {
				goodSuffixes[j] = j - i
			}

			j = borderPos[j]
		}

		i--
		j--
		borderPos[i] = j
	}

	j = borderPos[0]
	for i = 0; i <= patternLen; i++ {
		if goodSuffixes[i] == 0 {
			goodSuffixes[i] = j
		}

		if i == j {
			j = borderPos[j]
		}
	}

	return goodSuffixes
}

// FindBC finds all occurrences of substring "pattern"
// in a string "text" using the Boyer-Moore method (bad char shifts).
func FindBC(text, pattern string) []int {
	textLen := len(text)
	patternLen := len(pattern)
	badCharShifts := getBadCharShifts(pattern)
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

		shift := j - badCharShifts[text[i+j]]
		i += int(math.Max(1, float64(shift)))
	}

	return entries
}

// FindGS finds all occurrences of substring "pattern"
// in a string "text" using the Boyer-Moore method (good suffixes shifts).
func FindGS(text, pattern string) []int {
	textLen := len(text)
	patternLen := len(pattern)
	goodSuffixShifts := getGoodSuffixShifts(pattern)
	entries := make([]int, 0)

	i := 0
	for i <= textLen-patternLen {
		j := patternLen - 1

		for j >= 0 && text[i+j] == pattern[j] {
			j--
		}

		if j < 0 {
			entries = append(entries, i)
			i += goodSuffixShifts[0]
		} else {
			shift := goodSuffixShifts[j+1]
			i += int(math.Max(1, float64(shift)))
		}
	}

	return entries
}
