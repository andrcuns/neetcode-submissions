func anagramMappings(nums1 []int, nums2 []int) []int {
	positions := make(map[int]int)

	for i, num := range nums2 {
		positions[num] = i
	}

	res := make([]int, 0, len(nums1))

	for _, num := range nums1 {
		res = append(res, positions[num])
	}

	return res
}
