package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(tableFib(n))
}

func tableFib(n int) uint {
	var numbers = make([]uint, n+1)

	for i := range numbers {
		if i <= 1 {
			numbers[i] = uint(i)
		} else {
			numbers[i] = numbers[i-1] + numbers[i-2]
		}
	}

	return numbers[n]
}
