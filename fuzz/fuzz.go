package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(fuzz_brute(10))
}

// with if else
func fuzz_brute(n int) []string {
	var a []string

	for i := 0; i < n; i++ {
		str := ""

		if i % 3 == 0 { 
			str += "Fizz" 
		} else if i % 5 == 0 { 
			str += "Buzz" 
		} else { str = strconv.Itoa(i) }

		a = append(a, str)
	}

	return a
}

// with switch case (and more elegant checking)
func fuzz_switch(n int) []string {
	out := make([]string, n)

	for i := 0; i < n; i++ {
		switch (i + 1) % 15 {
			case 0:
				out[i] = "FizzBuzz"
			case 3, 6, 9, 12:
				out[i] = "Fizz"
			case 5, 10:
				out[i] = "Buzz"
			default:
				out[i] = strconv.Itoa(i + 1)
		}
	}

	return out
}