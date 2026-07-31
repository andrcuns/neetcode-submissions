func hasDuplicate(nums []int) bool {
    keys := make(map[int]bool)

	for _, num := range nums {
		if keys[num] {
			return true
		}
		keys[num] = true
	}
	return false
}
