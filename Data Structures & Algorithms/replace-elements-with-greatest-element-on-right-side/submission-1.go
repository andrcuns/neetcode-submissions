func replaceElements(arr []int) []int {
	n := len(arr)
	ans := make([]int, n)
	max := -1

	for i := n-1; i >=0; i-- {
		ans[i] = max
		if arr[i] > max {
			max = arr[i]
		}
	}
	return ans
}
