func isValid(s string) bool {
    stack := make([]rune, 0, len(s))
	mapping := map[rune]rune{')': '(', '}': '{', ']': '['}
	for _, char := range s {
		if open, ok := mapping[char]; ok {
			if len(stack) != 0 {
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if top != open {
					return false
				}
			} else {
				return false
			}
		} else {
			stack = append(stack, char)
		}
	}

	return len(stack) == 0
}
