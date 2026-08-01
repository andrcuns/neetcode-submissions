func search(nums []int, target int) int {
	return binarySearch(0, len(nums)-1, nums, target)
}

func binarySearch(start int, finish int, nums []int, target int) int {
	if start > finish {
		return -1
	}

	middle := start + (finish-start)/2

	if nums[middle] == target {
		return middle
	}

	if nums[middle] > target {
		return binarySearch(start, middle-1, nums, target)
	}

	return binarySearch(middle+1, finish, nums, target)
}
