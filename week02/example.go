package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	numbers := tableFib(n)

	fmt.Println(numbers)

	for _, i := range numbers {
		fmt.Printf("%d ", i%2)
	}
	fmt.Printf("\n")

	for _, i := range numbers {
		fmt.Printf("%d ", i%3)
	}
	fmt.Printf("\n")

	for _, i := range numbers {
		fmt.Printf("%d ", i%4)
	}
	fmt.Printf("\n")
}

func tableFib(n int) []uint {
	var numbers = make([]uint, n+1)

	for i := range numbers {
		if i <= 1 {
			numbers[i] = uint(i)
		} else {
			numbers[i] = numbers[i-1] + numbers[i-2]
		}
	}

	return numbers
}
