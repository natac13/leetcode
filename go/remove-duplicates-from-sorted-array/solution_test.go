package removeduplicatesfromsortedarray

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "example 1",
			nums: []int{1, 1, 2},
			want: 2,
		},
		{
			name: "example 2",
			nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			want: 5,
		},
		// {
		// 	name: "empty input",
		// 	nums: []int{},
		// 	want: 0,
		// },
		// {
		// 	name: "single element",
		// 	nums: []int{1},
		// 	want: 1,
		// },
		// {
		// 	name: "all duplicates",
		// 	nums: []int{1, 1, 1, 1, 1},
		// 	want: 1,
		// },
		// {
		// 	name: "no duplicates",
		// 	nums: []int{1, 2, 3, 4, 5},
		// 	want: 5,
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeDuplicates(tt.nums)
			if got != tt.want {
				t.Errorf("removeDuplicates(%v) = %v, want %v", tt.nums, got, tt.want)
			}
			if !reflect.DeepEqual(tt.nums[:tt.want], tt.nums[:got]) {
				t.Errorf("removeDuplicates(%v) did not correctly modify the input slice", tt.nums)
			}
		})
	}
}
