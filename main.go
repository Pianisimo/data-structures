/*
Just a playground
*/

package main

import "fmt"

func main() {
	//s := []int{1, 2, 3, 4, 5}
	//fmt.Println(s)
	//
	//s = s[1:]
	//fmt.Println(s)

	//f, err := calculator.Calculate("2 /2+3 * 4.75- -(6)")
	//if err != nil {
	//	return
	//}
	//fmt.Println(f)
	x := []int{1, 2, 3, 4, 5, 6}

	fmt.Println(x)
	moveToStart(x, 5)
	fmt.Println(x)
}

func moveToStart(sl []int, value int) {
	for i := 0; i < len(sl); i++ {
		if sl[i] == value {
			for j := i; j > 0; j-- {
				tmp := sl[j]
				sl[j] = sl[j-1]
				sl[j-1] = tmp
			}
		}
	}
}
