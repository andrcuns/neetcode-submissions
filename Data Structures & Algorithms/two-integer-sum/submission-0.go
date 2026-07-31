func twoSum(nums []int, target int) []int {
    prevVal := make(map[int]int)

	for i, num := range nums {
		diff := target - num
		j, found := prevVal[diff]

		if found {
			return []int{j, i}
		}
        prevVal[num] = i
	}
	return []int{}
}
