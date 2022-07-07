package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(lcm(a, b))
}

// lcm computes the least common multiple of two positive numbers
func lcm(a, b int) int {
	gcd := gcd(a, b)
	return (a * b) / gcd
}

// gcd Computes the greatest common divisor between two numbers
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	r := a % b
	return gcd(b, r)
}
