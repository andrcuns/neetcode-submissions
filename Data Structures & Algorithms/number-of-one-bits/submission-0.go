func hammingWeight(n int) int {
	res := 0
	for _, char := range fmt.Sprintf("%b", n) {
		if char == '1' {
			res++
		}
	}

	return res
}
