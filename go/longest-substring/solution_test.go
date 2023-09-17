// BEGIN: 3f1c8c7f9b5d
package longestsubstring

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"", 0},
		{" ", 1},
		{"au", 2},
		{"aab", 2},
		{"dvdf", 3},
	}

	for _, tc := range testCases {
		result := lengthOfLongestSubstring(tc.input)
		if result != tc.expected {
			t.Errorf("For input %q, expected %d but got %d", tc.input, tc.expected, result)
		}
	}
}

// END: 3f1c8c7f9b5d
