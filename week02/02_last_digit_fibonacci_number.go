package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(tableFib(n))
}

func tableFib(n int) int {
	var numbers = make([]int, n+1)

	for i := range numbers {
		if i <= 1 {
			numbers[i] = int(i)
		} else {
			numbers[i] = (numbers[i-1] + numbers[i-2]) % 10
		}
	}

	return numbers[n]
}
