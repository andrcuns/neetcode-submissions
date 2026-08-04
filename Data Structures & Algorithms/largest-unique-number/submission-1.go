func largestUniqueNumber(nums []int) int {
	mapping := make(map[int]int)
	
	for _, num := range nums {
		mapping[num]++
	}

	max := -1

	for k, v := range mapping {
		if k > max && v < 2 {
			max = k
		}
	}

	return max
}
