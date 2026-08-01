func isPalindrome(s string) bool {
    chars := strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			return unicode.ToLower(r)
		}
		return -1
	}, s)

	for i, char := range chars {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			if chars[i] != chars[len(chars)-(i+1)] {
				return false
			}
		}
	}

	return true
}
