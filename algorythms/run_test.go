package main

import (
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
