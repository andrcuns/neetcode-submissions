func singleNumber(nums []int) int {
	numMap := make(map[int]int)

	for _, num := range nums {
		numMap[num]++
	}

	for key, val := range numMap {
		if val == 1 {
			return key
		}
	}

	return -1
}
