package main

var calculations = 0

func fibonachi(n int) int {
	calculations++
	if n < 2 {
		return n
	}

	return fibonachi(n-1) + fibonachi(n-2)
}

/*
Dynamic Programing / Caching
*/

func fibonachiCache() func(n int) int {
	cache := make(map[int]int, 0)

	var fib func(n int) int
	fib = func(n int) int {
		if cache[n] != 0 {
			return cache[n]
		} else {
			calculations++
			if n < 2 {
				return n
			}

			cache[n] = fib(n-1) + fib(n-2)
			return cache[n]
		}
	}
	return fib
}
