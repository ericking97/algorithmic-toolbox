package main

import (
	"fmt"
	"reflect"
)

func main() {
	var k, m uint
	fmt.Scan(&k, &m)
	fmt.Println(ComputeFibMod(k, m))
}

// ComputeFibMod Computes the given mod m for the k'th fibonacci number
func ComputeFibMod(k, m uint) uint {
	if k <= 1 {
		return uint(k)
	}

	var nextFib uint
	fibSequence := []uint{0, 1}
	pisanoPeriod := []uint{}
	hasPeriod := false

	j := 0
	for i := uint(2); i <= k; i++ {
		nextFib = (fibSequence[i-1] + fibSequence[i-2]) % m
		fibSequence = append(fibSequence, nextFib)

		if nextFib == fibSequence[j] {
			pisanoPeriod = append(pisanoPeriod, nextFib)
			j++
			if reflect.DeepEqual(pisanoPeriod, fibSequence[0:len(fibSequence)-len(pisanoPeriod)]) {
				hasPeriod = true
				break
			}
		} else {
			pisanoPeriod = []uint{}
			j = 0
		}
	}

	if hasPeriod {
		pos := int(k) % len(pisanoPeriod)
		kthfib := pisanoPeriod[pos]
		return kthfib
	}

	return nextFib
}
