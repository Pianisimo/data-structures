package main

import "fmt"

func main() {
	customHashMap := NewCustomHashMap(2)

	customHashMap.Set("key1", 1000)
	customHashMap.Set("key2", 234)
	customHashMap.Set("key3", 3333)
	fmt.Printf("key1: %v\n", customHashMap.Get("key1"))
	fmt.Printf("key2: %v\n", customHashMap.Get("key2"))
	fmt.Printf("key3: %v\n", customHashMap.Get("key3"))
	fmt.Printf("key4: %v\n", customHashMap.Get("key4"))
	fmt.Printf("customHashMap: %v\n", customHashMap)
	fmt.Printf("keys: %v\n", customHashMap.GetKeys())

	s := []int{1, 0, 0, 4, 4, 5, 0, 6}
	fmt.Printf("first duplicate: %v\n", findFirstDuplicate(s))
}

func findFirstDuplicate(s []int) int {
	m := map[int]bool{}

	for i := 0; i < len(s); i++ {
		if !m[s[i]] {
			m[s[i]] = true
		} else {
			return s[i]
		}
	}

	return -1
}
