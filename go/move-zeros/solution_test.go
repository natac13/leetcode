package movezeros

import (
	"reflect"
	"testing"
)

func TestMoveZeros(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "Test 1 - starts with zero",
			nums:     []int{0, 1, 0, 3, 12},
			expected: []int{1, 3, 12, 0, 0},
		},
		{
			name:     "Test 2 - No Zeros",
			nums:     []int{2, 1},
			expected: []int{2, 1},
		},
		{
			name:     "Test 3 - All zeros",
			nums:     []int{0, 0, 0, 0},
			expected: []int{0, 0, 0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeros(tt.nums)
			if !reflect.DeepEqual(tt.nums, tt.expected) {
				t.Errorf("moveZeros() = %v, wanted %v", tt.nums, tt.expected)
			}
		})
	}
}

func TestMoveZerosV2(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "Test 1 - starts with zero",
			nums:     []int{0, 1, 0, 3, 12},
			expected: []int{1, 3, 12, 0, 0},
		},
		{
			name:     "Test 2 - No Zeros",
			nums:     []int{2, 1},
			expected: []int{2, 1},
		},
		{
			name:     "Test 3 - All zeros",
			nums:     []int{0, 0, 0, 0},
			expected: []int{0, 0, 0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZerosV2(tt.nums)
			if !reflect.DeepEqual(tt.nums, tt.expected) {
				t.Errorf("moveZerosV2() = %v, wanted %v", tt.nums, tt.expected)
			}
		})
	}
}
