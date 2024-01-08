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
	var r, s int

	r = p % q
	s = r 

	for s > 0 {
		s = q % r
		if (s != 0) { r = s }
	}
	
	if (r == 0) { return q } else { return r }
}