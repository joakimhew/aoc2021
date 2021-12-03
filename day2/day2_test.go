package main

import "testing"

// TestCalculateSubmarinePositionPart1
func TestCalculateSubmarinePositionPart1(t *testing.T) {
	testCases := []struct {
		movements []movement
		expected  int
	}{
		{[]movement{{forward, 5}, {down, 5}, {forward, 8}, {up, 3}, {down, 8}, {forward, 2}}, 150},
	}

	for _, testCase := range testCases {
		actual := calculateSubmarinePositionPart1(testCase.movements)
		if actual != testCase.expected {
			t.Errorf("Expected %d, got %d", testCase.expected, actual)
		}
	}
}

// TestCalculateSubmarinePositionPart2
func TestCalculateSubmarinePositionPart2(t *testing.T) {
	testCases := []struct {
		movements []movement
		expected  int
	}{
		{[]movement{{forward, 5}, {down, 5}, {forward, 8}, {up, 3}, {down, 8}, {forward, 2}}, 900},
	}

	for _, testCase := range testCases {
		actual := calculateSubmarinePositionPart2(testCase.movements)
		if actual != testCase.expected {
			t.Errorf("Expected %d, got %d", testCase.expected, actual)
		}
	}
}
