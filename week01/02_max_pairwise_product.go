package main

import "fmt"

// MaxPairProd Computes the greatest value obtained by multiplying the two
// greatest numbers in a slice
func MaxPairProd(numbers []uint) uint {
	var i, j uint

	for _, n := range numbers {
		if n > i {
			j = i
			i = n
		} else {
			if n > j {
				j = n
			}
		}
	}

	return i * j
}

func main() {
	var size int
	fmt.Scan(&size)

	var numbers = make([]uint, size)
	for i := 0; i < size; i++ {
		fmt.Scan(&numbers[i])
	}

	fmt.Println(MaxPairProd(numbers))
}
