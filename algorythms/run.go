package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{2, 7, 15, 11}
	target := 9

	fmt.Println(two_sum_pointers(numbers, target))
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
		if s != 0 {
			r = s
		}
	}

	if r == 0 {
		return q
	} else {
		return r
	}
}

// brute force solution (worst)
func two_sum_brute(nums []int, target int) []int {
	var result []int

	n := len(nums)

	for i := 0; i < n - 1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result = []int{i, j}
				break
			}
		}
	}

	return result
}

// hashmap solution
func two_sum_hashmap(nums []int, target int) []int {
	num_map := make(map[int]int)

	for i, num := range nums {
		complement := target - num

		if index, ok := num_map[complement]; ok {
			return []int{index, i}
		}
		num_map[num] = i
	}

	return []int{}
}

// 2 pointers solution (best)
func two_sum_pointers(nums []int, target int) []int {
	sorted := make([][2]int, len(nums))

	for i, num := range(nums) {
		sorted[i] = [2]int{i, num}
	}

	sort.Slice(sorted, func (i, j int) bool {
		return sorted[i][1] < sorted[j][1]
	})

	left, right := 0, len(nums) - 1

	for left < right {
		sum := sorted[left][1] + sorted[right][1]

		if sum == target {
			return []int{sorted[left][0], sorted[right][0]}
		} else if sum > target {
			right--
		} else {
			left++
		}
	}

	return []int{}
}

