func isAnagram(s string, t string) bool {
    firstString := make(map[string]int)
	secondString := make(map[string]int)

	for _, char := range s {
		firstString[string(char)] += 1
	}
	for _, char := range t {
		secondString[string(char)] += 1
	}

	if len(firstString) != len(secondString) {
		return false
	}

	for key, value := range firstString {
		if secondString[key] != value {
			return false
		}
	}

	return true
}
