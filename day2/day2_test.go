package main

import "testing"

// TestCalculateSubmarinePosition
func TestCalculateSubmarinePosition(t *testing.T) {
	testCases := []struct {
		movements []movement
		expected  int
	}{
		{[]movement{{forward, 5}, {down, 5}, {forward, 8}, {up, 3}, {down, 8}, {forward, 2}}, 150},
	}

	for _, testCase := range testCases {
		actual := calculateSubmarinePosition(testCase.movements)
		if actual != testCase.expected {
			t.Errorf("Expected %d, got %d", testCase.expected, actual)
		}
	}
}
