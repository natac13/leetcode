package removeelement

import (
	"reflect"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		val      int
		expected int
		wantNums []int
	}{
		{
			name:     "example 1",
			nums:     []int{3, 2, 2, 3},
			val:      3,
			expected: 2,
			wantNums: []int{2, 2},
		},
		{
			name:     "example 2",
			nums:     []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:      2,
			expected: 5,
			wantNums: []int{0, 1, 3, 0, 4},
		},
		{
			name:     "empty slice",
			nums:     []int{},
			val:      1,
			expected: 0,
			wantNums: []int{},
		},
		{
			name:     "no elements to remove",
			nums:     []int{1, 2, 3},
			val:      4,
			expected: 3,
			wantNums: []int{1, 2, 3},
		},
		{
			name:     "remove all elements",
			nums:     []int{1, 1, 1, 1},
			val:      1,
			expected: 0,
			wantNums: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeElement(tt.nums, tt.val)
			if got != tt.expected {
				t.Errorf("removeElement() = %v, want %v", got, tt.expected)
			}
			if !reflect.DeepEqual(tt.nums[:got], tt.wantNums) && !slicesEqualIgnoreOrder(tt.nums[:got], tt.wantNums) {
				t.Errorf("removeElement() modified nums incorrectly, got %v, want %v, where the full slice is %v", tt.nums[:got], tt.wantNums, tt.nums)
			}
		})
	}
}

func slicesEqualIgnoreOrder(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	m := make(map[int]int)
	for _, x := range a {
		m[x]++
	}
	for _, x := range b {
		if m[x] == 0 {
			return false
		}
		m[x]--
	}
	return true
}

func TestRemoveElementV2(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		val      int
		expected int
		wantNums []int
	}{
		{
			name:     "example 1",
			nums:     []int{3, 2, 2, 3},
			val:      3,
			expected: 2,
			wantNums: []int{2, 2},
		},
		{
			name:     "example 2",
			nums:     []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:      2,
			expected: 5,
			wantNums: []int{0, 1, 3, 0, 4},
		},
		{
			name:     "empty slice",
			nums:     []int{},
			val:      1,
			expected: 0,
			wantNums: []int{},
		},
		{
			name:     "no elements to remove",
			nums:     []int{1, 2, 3},
			val:      4,
			expected: 3,
			wantNums: []int{1, 2, 3},
		},
		{
			name:     "remove all elements",
			nums:     []int{1, 1, 1, 1},
			val:      1,
			expected: 0,
			wantNums: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeElementV2(tt.nums, tt.val)
			if got != tt.expected {
				t.Errorf("removeElementV2() = %v, want %v", got, tt.expected)
			}
			if !reflect.DeepEqual(tt.nums[:got], tt.wantNums) && !slicesEqualIgnoreOrder(tt.nums[:got], tt.wantNums) {
				t.Errorf("removeElementV2() modified nums incorrectly, got %v, want %v, where the full slice is %v", tt.nums[:got], tt.wantNums, tt.nums)
			}
		})
	}
}

// benchmarks

func BenchmarkRemoveElement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		removeElement([]int{3, 2, 2, 3}, 3)
	}
}

func BenchmarkRemoveElementV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		removeElementV2([]int{3, 2, 2, 3}, 3)
	}
}
