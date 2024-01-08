package main

import (
	"reflect"
	"testing"
)

func TestEuclideanRecursive(t *testing.T) {
	tests := []struct {
		p        int
		q        int
		expected int
	}{
		{48, 18, 6},
		{35, 14, 7},
		{81, 27, 27},
		{100, 75, 25},
		{17, 5, 1},
		{78, 66, 6},
		{81, 27, 27},
	}

	for _, test := range tests {
		result := euclidean_recursive(test.p, test.q)
		if result != test.expected {
			t.Errorf("For (%d, %d), expected %d, but got %d", test.p, test.q, test.expected, result)
		}
	}
}

func TestEuclideanLoop(t *testing.T) {
	tests := []struct {
		p        int
		q        int
		expected int
	}{
		{48, 18, 6},
		{35, 14, 7},
		{81, 27, 27},
		{100, 75, 25},
		{17, 5, 1},
		// Add more test cases as needed
	}

	for _, test := range tests {
		result := euclidean_loop(test.p, test.q)
		if result != test.expected {
			t.Errorf("For (%d, %d), expected %d, but got %d", test.p, test.q, test.expected, result)
		}
	}
}

func TestTwoSumBrute(t *testing.T) {
	// Test case 1: Valid input, expect [1, 2]
	nums1 := []int{2, 7, 11, 15}
	target1 := 18
	expected1 := []int{1, 2}
	result1 := two_sum_brute(nums1, target1)
	if !reflect.DeepEqual(result1, expected1) {
		t.Errorf("Test case 1 failed. Expected %v, got %v", expected1, result1)
	}

	// Test case 2: No solution, expect []
	nums2 := []int{2, 7, 11, 15}
	target2 := 3
	expected2 := []int{}
	result2 := two_sum_brute(nums2, target2)
	if !reflect.DeepEqual(result2, expected2) {
		t.Errorf("Test case 2 failed. Expected %v, got %v", expected2, result2)
	}

	// Test case 3: Empty input, expect []
	nums3 := []int{}
	target3 := 5
	expected3 := []int{}
	result3 := two_sum_brute(nums3, target3)
	if !reflect.DeepEqual(result3, expected3) {
		t.Errorf("Test case 3 failed. Expected %v, got %v", expected3, result3)
	}
}

func TestTwoSumHashmap(t *testing.T) {
	// Test case 1: Valid input, expect [1, 2]
	nums1 := []int{2, 7, 11, 15}
	target1 := 18
	expected1 := []int{1, 2}
	result1 := two_sum_hashmap(nums1, target1)
	if !reflect.DeepEqual(result1, expected1) {
		t.Errorf("Test case 1 failed. Expected %v, got %v", expected1, result1)
	}

	// Test case 2: No solution, expect []
	nums2 := []int{2, 7, 11, 15}
	target2 := 3
	expected2 := []int{}
	result2 := two_sum_hashmap(nums2, target2)
	if !reflect.DeepEqual(result2, expected2) {
		t.Errorf("Test case 2 failed. Expected %v, got %v", expected2, result2)
	}

	// Test case 3: Empty input, expect []
	nums3 := []int{}
	target3 := 5
	expected3 := []int{}
	result3 := two_sum_hashmap(nums3, target3)
	if !reflect.DeepEqual(result3, expected3) {
		t.Errorf("Test case 3 failed. Expected %v, got %v", expected3, result3)
	}
}

func TestTwoSumPointers(t *testing.T) {
	// Test case 1: Valid input, expect [1, 2]
	nums1 := []int{2, 7, 11, 15}
	target1 := 18
	expected1 := []int{1, 2}
	result1 := two_sum_pointers(nums1, target1)
	if !reflect.DeepEqual(result1, expected1) {
		t.Errorf("Test case 1 failed. Expected %v, got %v", expected1, result1)
	}

	// Test case 2: No solution, expect []
	nums2 := []int{2, 7, 11, 15}
	target2 := 3
	expected2 := []int{}
	result2 := two_sum_pointers(nums2, target2)
	if !reflect.DeepEqual(result2, expected2) {
		t.Errorf("Test case 2 failed. Expected %v, got %v", expected2, result2)
	}

	// Test case 3: Empty input, expect []
	nums3 := []int{}
	target3 := 5
	expected3 := []int{}
	result3 := two_sum_pointers(nums3, target3)
	if !reflect.DeepEqual(result3, expected3) {
		t.Errorf("Test case 3 failed. Expected %v, got %v", expected3, result3)
	}
}