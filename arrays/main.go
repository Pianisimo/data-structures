package main

import "fmt"

/*
Not many methods to work with in arrays with go.
*/

func main() {
	var array1 [5]int
	array1[0] = 10
	array2 := [5]int{0: 10}
	fmt.Println("array1:", array1)

	if array1 == array2 {
		fmt.Println("matches")
	} else {
		fmt.Println("does not matches")
	}

	slice := []int{0: 10, 1: 2, 2: 3, 3: 4, 4: 5}
	reversed := reverseSlice(slice)
	fmt.Println("slice:", slice)
	fmt.Println("reversed:", reversed)

	sorted1 := []int{1, 2, 7, 8, 9, 12, 13, 14}
	sorted2 := []int{3, 4, 5, 6, 10, 11}
	merged := mergeOrderedSlices(sorted1, sorted2)
	fmt.Println("merged:", merged)
}

/*
Makes no sense with arrays
*/
func reverseSlice[T any](slice []T) []T {
	var backwards = make([]T, len(slice))
	for i, value := range slice {
		backwards[len(slice)-i-1] = value
	}

	return backwards
}

func mergeOrderedSlices(slice1 []int, slice2 []int) []int {
	var ordered []int
	if len(slice1) == 0 {
		return slice1
	}
	if len(slice2) == 0 {
		return slice2
	}

	item1 := slice1[0]
	item2 := slice2[0]
	i := 1
	j := 1

	for true {
		if item1 < item2 {
			ordered = append(ordered, item1)
			item1 = slice1[i]
			i++
		} else {
			ordered = append(ordered, item2)
			item2 = slice2[j]
			j++
		}

		if i == len(slice1) || j == len(slice2) {
			if i < len(slice1) {
				ordered = append(ordered, item2)
				ordered = append(ordered, slice1[i:]...)
			} else {
				ordered = append(ordered, item1)
				ordered = append(ordered, slice2[j:]...)
			}
			break
		}
	}

	return ordered
}
