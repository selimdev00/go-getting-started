package main

import (
	"fmt"
)

func main() {
	var p, q int

	fmt.Println("Type numbers: ")
	fmt.Scan(&p, &q)

	fmt.Printf("The result is: %v", euclidean(p, q))
}

// recursive algorythm, getting remainder and starting same functio with q and remainder
func euclidean(p int, q int) int {
	r := p % q

	if r != 0 {
		return euclidean(q, r)
	}
	
	return q
}