package algorithms

func InsertionSort(s []int) {
	for i := 0; i < len(s); i++ {
		if s[i] < s[0] {
			temp := s[0]
			s[0] = s[i]
			s[i] = temp
		} else {
			for j := 1; j < i; j++ {
				if s[i] > s[j-1] && s[i] < s[j] {
					temp := s[j]
					s[j] = s[i]
					s[i] = temp
				}
			}
		}
	}
}
