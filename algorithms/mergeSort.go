package algorithms

func MergeSort(s []int) {
	copy(s, recursiveMergeSort(s))
}

func recursiveMergeSort(s []int) []int {
	if len(s) == 1 {
		return s
	}

	var (
		length = len(s)
		middle = length / 2
		left   = s[:middle]
		right  = s[middle:]
	)

	s = merge(recursiveMergeSort(left), recursiveMergeSort(right))
	return s
}

func merge(left, right []int) []int {
	result := make([]int, 0)
	leftIndex := 0
	rightIndex := 0

	for leftIndex < len(left) && rightIndex < len(right) {
		if left[leftIndex] < right[rightIndex] {
			result = append(result, left[leftIndex])
			leftIndex++
		} else {
			result = append(result, right[rightIndex])
			rightIndex++
		}
	}

	result = append(result, left[leftIndex:]...)
	result = append(result, right[rightIndex:]...)
	return result
}
