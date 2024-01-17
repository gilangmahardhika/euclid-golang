package main

import "fmt"

func gcd(p int, q int) int {
	if q == 0 {
		return p
	} else {
		r := p % q
		return r
	}
}

func main() {
	gcd := gcd(124, 15)
	fmt.Println(gcd)
}
