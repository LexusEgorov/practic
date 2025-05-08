package main

import "testing"

type DigitTest struct {
	Input    int
	Expected int
}

func TestSumDigits(t *testing.T) {
	tests := []DigitTest{
		{
			Input:    0,
			Expected: 0,
		},
		{
			Input:    123,
			Expected: 6,
		},
		{
			Input:    9876543210,
			Expected: 45,
		},
		{
			Input:    111,
			Expected: 3,
		},
		{
			Input:    -123,
			Expected: 6,
		},
	}

	for i, test := range tests {
		res := SumDigits(test.Input)

		if res != test.Expected {
			t.Errorf("Error in test №%d, got: %d, want: %d", i+1, res, test.Expected)
		}
	}
}
