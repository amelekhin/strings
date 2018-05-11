package rabinkarp

const q = 861
const charNumber = 256

func getPatternHash(pattern string) int {
	patternHash := 0

	for i := range pattern {
		patternHash = (charNumber*patternHash + int(pattern[i])) % q
	}

	return patternHash
}

func getH(pattern string) int {
	h := 1
	patternLen := len(pattern)

	for i := 0; i < patternLen-1; i++ {
		h = (h * charNumber) % q
	}

	return h
}

// Find finds all occurrences of substring "pattern"
// in a string "text" using the Rabin-Karp method.
func Find(text, pattern string) []int {
	entries := make([]int, 0)
	textLen := len(text)
	patternLen := len(pattern)
	patternHash := getPatternHash(pattern)
	h := getH(pattern)
	textHash := 0

	for i := range pattern {
		textHash = (charNumber*textHash + int(text[i])) % q
	}

	for i := 0; i <= textLen-patternLen; i++ {
		if patternHash == textHash {
			equal := true
			j := 0

			for equal && j < patternLen {
				if text[i+j] != pattern[j] {
					equal = false
				}

				j++
			}

			if equal {
				entries = append(entries, i)
			}
		}

		if i < textLen-patternLen {
			textHash = (charNumber*(textHash-int(text[i])*h) + int(text[i+patternLen])) % q

			if textHash < 0 {
				textHash += q
			}
		}
	}

	return entries
}
