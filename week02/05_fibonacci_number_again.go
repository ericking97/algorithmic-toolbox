package main

import (
	"fmt"
	"reflect"
)

func main() {
	var k, m int
	fmt.Scan(&k, &m)
	fmt.Println(ComputeFib(k, m))
}

// ComputeFibMod Computes the given mod m for the k'th fibonacci number
func ComputeFib(k, m int) uint {
	fibonacciNumbers := []uint{0, 1}
	var pisanoPeriod []uint
	var nextFib uint
	hasPeriod := false
	
	if k <= 1 {
		return uint(k)
	}

	j := 0
	for i := 2; i <= k; i++ {
		nextFib = (fibonacciNumbers[i-1] + fibonacciNumbers[i-2]) % uint(m)
		fibonacciNumbers = append(fibonacciNumbers, nextFib)

		if nextFib == fibonacciNumbers[j] {
			pisanoPeriod = append(pisanoPeriod, nextFib)
			j++
			if reflect.DeepEqual(pisanoPeriod, fibonacciNumbers[0:len(fibonacciNumbers)-len(pisanoPeriod)]) {
				hasPeriod = true
				break
			}
		} else {
			pisanoPeriod = []uint{}
			j = 0
		}
	}

	var fib uint
	if hasPeriod {
		position := k % len(pisanoPeriod)
		fib = pisanoPeriod[position]
	} else {
		fib = nextFib
	}

	return fib
}
