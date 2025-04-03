package main

import "testing"

func TestSubarraysWithKDistinct(t *testing.T) {
	cases := []struct {
		input []int
		k     int
		want  int
	}{
		{
			[]int{1, 2, 1, 2, 3},
			2,
			7,
		},
		{
			[]int{1, 2, 1, 3, 4},
			3,
			3,
		},
	}

	for _, tc := range cases {
		got := subarraysWithKDistinct(tc.input, tc.k)
		if got != tc.want {
			t.Errorf("subarraysWithKDistinct(%v, %d) = %d; want %d", tc.input, tc.k, got, tc.want)
		}
	}
}
