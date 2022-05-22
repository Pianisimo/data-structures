package main

func bubblesSort(s []int) {
	for j := 0; j < len(s); j++ {
		for i := 0; i < len(s)-1; i++ {
			if s[i] > s[i+1] {
				temp := s[i]
				s[i] = s[i+1]
				s[i+1] = temp
			}
		}
	}
}
