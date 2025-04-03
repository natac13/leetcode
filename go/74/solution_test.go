package main

import "testing"

func TestSearchMatrix(t *testing.T) {
	cases := []struct {
		matrix [][]int
		target int
		want   bool
	}{
		{
			matrix: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 50},
			},
			target: 3,
			want:   true,
		},
		{
			matrix: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 50},
			},
			target: 13,
			want:   false,
		},
	}

	for _, tc := range cases {
		got := searchMatrix(tc.matrix, tc.target)
		if got != tc.want {
			t.Errorf("searchMatrix(%v, %d) = %v, want %v", tc.matrix, tc.target, got, tc.want)
		}
	}
}
