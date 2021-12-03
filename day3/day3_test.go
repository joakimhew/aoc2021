package main

import "testing"

// testCalculateGammaAndEpsRatePart1
func TestCalculateGammaAndEpsRatePart1(t *testing.T) {
	testCases := []struct {
		data     []string
		expected int64
	}{
		{
			data: []string{
				"00100",
				"11110",
				"10110",
				"10111",
				"10101",
				"01111",
				"00111",
				"11100",
				"10000",
				"11001",
				"00010",
				"01010",
			},
			expected: 198,
		},
	}

	for _, tc := range testCases {
		actual, err := calcPart1(tc.data, 5)
		if err != nil {
			t.Errorf("Error: %v", err)
		}
		if actual != tc.expected {
			t.Errorf("Expected: %v, Actual: %v", tc.expected, actual)
		}
	}
}
