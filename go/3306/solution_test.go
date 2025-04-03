package main

import "testing"

func TestCountOfSubstrings(t *testing.T) {
	cases := []struct {
		word string
		k    int
		want int64
	}{
		{"aeioqq", 1, 0},
		{"aeiou", 0, 1},
		{"ieaouqqieaouqq", 1, 3},
		{"eoeaui", 0, 2},
	}

	for _, c := range cases {
		got := countOfSubstrings(c.word, c.k)
		if got != c.want {
			t.Errorf("countOfSubstrings(%q, %d) = %d; want %d", c.word, c.k, got, c.want)
		}
	}
}
