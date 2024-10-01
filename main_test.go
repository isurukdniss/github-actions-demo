package main

import (
	"testing"
)

func TestCalculateSum(t *testing.T) {

	// Arrange
	tests := []struct {
		x        int
		y        int
		expected int
	}{
		{x: 1, y: 2, expected: 3},
		{x: -1, y: 2, expected: 1},
		{x: -1, y: -1, expected: -2},
	}

	for _, test := range tests {
		// Act
		result := calculateSum(test.x, test.y)

		// Assert
		if result != test.expected {
			t.Errorf("calculateSum(%d,%d) = %d; want %d", test.x, test.y, test.expected, result)
		}
	}
}
