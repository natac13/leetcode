package longestpalindromicsubstring

import "testing"

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "Example 1",
			s:    "babad",
			want: "bab",
		},
		{
			name: "Example 2",
			s:    "cbbd",
			want: "bb",
		},
		{
			name: "Single Character",
			s:    "a",
			want: "a",
		},
		{
			name: "Empty String",
			s:    "",
			want: "",
		},
		{
			name: "All Same Characters",
			s:    "ccc",
			want: "ccc",
		},
		{
			name: "Even Length Palindrome",
			s:    "abba",
			want: "abba",
		},
		{
			name: "Odd Length Palindrome",
			s:    "abcba",
			want: "abcba",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLongestPalindromeV2(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "Example 1",
			s:    "babad",
			want: "bab",
		},
		{
			name: "Example 2",
			s:    "cbbd",
			want: "bb",
		},
		{
			name: "Single Character",
			s:    "a",
			want: "a",
		},
		{
			name: "Empty String",
			s:    "",
			want: "",
		},
		{
			name: "All Same Characters",
			s:    "ccc",
			want: "ccc",
		},
		{
			name: "Even Length Palindrome",
			s:    "abba",
			want: "abba",
		},
		{
			name: "Odd Length Palindrome",
			s:    "abcba",
			want: "abcba",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindromeV2(tt.s); got != tt.want {
				t.Errorf("longestPalindromeV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
