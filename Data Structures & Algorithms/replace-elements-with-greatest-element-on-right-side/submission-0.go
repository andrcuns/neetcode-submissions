import "slices"

func replaceElements(arr []int) []int {
	for i := range arr {
		if i == len(arr)-1 {
			arr[i] = -1
			continue
		}
		
		highest := slices.Max(arr[i+1:])
		arr[i] = highest
	}

	return arr
}
