package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(gcd(a, b))
}

// gcd computes the greatest common divisor between two numbers
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	r := a % b
	return gcd(b, r)
}
