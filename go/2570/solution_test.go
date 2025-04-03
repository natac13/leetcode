package main

import (
	"reflect"
	"testing"
)

func TestMergeArrays(t *testing.T) {
	cases := []struct {
		nums1 [][]int
		nums2 [][]int
		want  [][]int
	}{
		{
			nums1: [][]int{
				{1, 2},
				{2, 3},
				{4, 5},
			},
			nums2: [][]int{
				{1, 4},
				{3, 2},
				{4, 1},
			},
			want: [][]int{
				{1, 6},
				{2, 3},
				{3, 2},
				{4, 6},
			},
		},
	}

	for _, tc := range cases {
		got := mergeArrays(tc.nums1, tc.nums2)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("mergeArrays(%v, %v) = %v, want %v", tc.nums1, tc.nums2, got, tc.want)
		}
	}
}
