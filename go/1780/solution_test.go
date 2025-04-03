package main

import "testing"

func TestCheckPowerOfThree(t *testing.T) {
	cases := []struct {
		input       int
		output      bool
		explanation string
	}{
		// {input: 12, output: true, explanation: "12 = 3^1 + 3^2"},
		// {input: 91, output: true, explanation: "91 = 3^0 + 3^2 + 3^4"},
		{input: 21, output: false, explanation: "given 21 return false"},
	}

	for _, tc := range cases {
		t.Run(tc.explanation, func(t *testing.T) {
			result := checkPowersOfThree(tc.input)
			if result != tc.output {
				t.Errorf("expected %v, got %v", tc.output, result)
			}
		})
	}
}
