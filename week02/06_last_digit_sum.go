package main

import (
	"fmt"
	"reflect"
)

func main() {
	var k uint
	fmt.Scan(&k)
	fmt.Println(ComputeSumLastDigit(k))
}

// ComputeSumLastDigit Computes the last digit of a sum for the nth fibonacci
// numbers
func ComputeSumLastDigit(k uint) uint {
	fibSequence := []uint{0, 1}
	total := uint(1)
	var pisanoPeriod []uint
	hasPeriod := false

	if k <= 1 {
		return k
	}

	j := 0
	for i := uint(2); i <= k; i++ {
		nextFib := (fibSequence[i-1] + fibSequence[i-2]) % 10
		fibSequence = append(fibSequence, nextFib)
		total = (total + nextFib) % 10
		if nextFib == fibSequence[j] {
			pisanoPeriod = append(pisanoPeriod, nextFib)
			j++
			if reflect.DeepEqual(pisanoPeriod, fibSequence[0:len(fibSequence) - len(pisanoPeriod)]) {
				hasPeriod = true
				break
			}
		} else {
			pisanoPeriod = []uint{}
			j = 0
		}
	}

	if hasPeriod {
		total = 0
		floor := k / uint(len(pisanoPeriod))
		pos := k % uint(len(pisanoPeriod))
		sumatory := uint(0)
		for _, v := range pisanoPeriod {
			sumatory += v
		}
		total = sumatory * floor

		for i := uint(0); i <= pos; i++ {
			total += pisanoPeriod[i]
		}
	}

	return total % 10
}

