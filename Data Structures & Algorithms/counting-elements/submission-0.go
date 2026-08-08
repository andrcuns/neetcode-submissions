func countElements(arr []int) int {
	mapping := make(map[int]bool)

	for _, num := range arr {
		mapping[num] = true
	}

	res := 0

	for _, num := range arr {
		if _, ok := mapping[num+1]; ok {
			res++
		}
	}

	return res
}
