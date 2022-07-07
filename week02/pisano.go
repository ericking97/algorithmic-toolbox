package main

import "fmt"

func main() {
	var kfib, mod int
	fmt.Scan(&kfib, &mod)
	Pisano(kfib, mod)
}

func Pisano(k, m int) {
	fibonacciSequence := []uint{0, 1}
	pisanoPeriod := []uint{0, 1}

	for i := 2; i <= k; i++ {
		nextFib := fibonacciSequence[i-1] + fibonacciSequence[i-2]
		fibonacciSequence = append(fibonacciSequence, nextFib)

		pisanoPeriod = append(pisanoPeriod, nextFib%uint(m))
	}

	fmt.Println(fibonacciSequence)
	fmt.Println(pisanoPeriod)
}
