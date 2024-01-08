package main

import (
	"fmt"
)

func main() {
	var p, q int

	fmt.Println("Type numbers: ")
	fmt.Scan(&p, &q)

	fmt.Printf("The result is: %v", euclidean_loop(p, q))
}

// recursive algorythm, getting remainder and starting same functio with q and remainder
func euclidean_recursive(p int, q int) int {
	r := p % q

	if r != 0 {
		return euclidean_recursive(q, r)
	}
	
	return q
}

// looping algorythm, without recursive
func euclidean_loop(p int, q int) int {
	r := p % q

	for r > 0 {
		r = q % r
	}
	
	return r
}