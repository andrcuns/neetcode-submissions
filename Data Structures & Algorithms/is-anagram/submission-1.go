func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
		return false
	}

	firstString := make(map[rune]int)
	secondString := make(map[rune]int)

	for i, char := range s {
		firstString[char]++
		secondString[rune(t[i])]++
	}

	for key, value := range firstString {
		if secondString[key] != value {
			return false
		}
	}

	return true
}
