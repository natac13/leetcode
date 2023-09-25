package validpalindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
		{"", true},
		{" ", true},
		{"  ", true},
		{"ab", false},
		{"aba", true},
		{"abba", true},
		{"abcba", true},
		{"abccba", true},
		{"abcdba", false},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			if isPalindrome(tc.input) != tc.expected {
				t.Errorf("expected %v but got %v", tc.expected, !tc.expected)
			}
		})
	}
}
