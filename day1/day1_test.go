package main

import (
	"testing"
)

// TestCalcPart1 is a test for the number of measurement increments
func TestCalcPart1(t *testing.T) {
	type test struct {
		input []int
		want  int
	}

	tests := []test{
		{[]int{1, 2, 2, 4, 3}, 2},
		{[]int{1, 2, 3, 4, 5}, 4},
	}

	for _, test := range tests {
		expectedNumMeasurementIncrements := test.want
		actualNumMeasurementIncrements := calcPart1(test.input)

		if expectedNumMeasurementIncrements != actualNumMeasurementIncrements {
			t.Errorf("Expected %d, but got %d", expectedNumMeasurementIncrements, actualNumMeasurementIncrements)
		}
	}
}

// TestCalcPart2 is a test for the number of measurement increments in a sliding window
func TestCalcPart2(t *testing.T) {
	type test struct {
		input []int
		want  int
	}

	tests := []test{
		{[]int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}, 5},
		{[]int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263, 280}, 6},
	}

	for _, test := range tests {
		expectedNumMeasurementIncrements := test.want
		actualNumMeasurementIncrements := calcPart2(test.input)

		if expectedNumMeasurementIncrements != actualNumMeasurementIncrements {
			t.Errorf("Expected %d, but got %d", expectedNumMeasurementIncrements, actualNumMeasurementIncrements)
		}
	}
}
