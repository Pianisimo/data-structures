package algorithms

func SelectionSort(s []int) {
	for j := 0; j < len(s); j++ {
		min := s[j]
		index := j

		for i := j + 1; i < len(s); i++ {
			if s[i] < min {
				min = s[i]
				index = i
			}
		}

		s[index] = s[j]
		s[j] = min
	}
}
