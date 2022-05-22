package main

import "fmt"

func main() {
	cache := fibonachiCache()
	for i := 0; i < 30; i++ {
		fmt.Printf("fibonachi of %v is: %v\n", i, fibonachi(i))
		fmt.Printf("calculations: %v\n", calculations)
		calculations = 0
		fmt.Println("****************************************")
		fmt.Printf("fibonachi of %v is: %v\n", i, cache(i))
		fmt.Printf("cached calculations: %v\n", calculations)
		fmt.Println("-----------------------------------------")
		calculations = 0
	}

	/*l := 20
	s := make([]int, l)

	for i := 0; i < l; i++ {
		s[i] = rand.Int() % 100

	}

	fmt.Printf("random: %v\n", s)
	mergeSort(s)
	fmt.Printf("sorted: %v\n", s)*/

}
