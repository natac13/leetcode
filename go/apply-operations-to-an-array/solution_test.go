package applyoperationstoanarray

import (
	"reflect"
	"testing"
)

func TestApplyOperations(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		expect []int
	}{
		{
			name:   "Example 1",
			nums:   []int{1, 2, 2, 1, 1, 0},
			expect: []int{1, 4, 2, 0, 0, 0},
		},
		{
			name:   "Example 2",
			nums:   []int{0, 1},
			expect: []int{1, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := applyOperations(tt.nums)
			if !reflect.DeepEqual(got, tt.expect) {
				t.Errorf("applyOperations() = %v, want %v", got, tt.expect)
			}
		})
	}
}

// benchmarks
func BenchmarkApplyOperationsV0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		applyOperationsV0([]int{1, 2, 2, 1, 1, 0})
	}
}

func BenchmarkApplyOperations(b *testing.B) {
	for i := 0; i < b.N; i++ {
		applyOperations([]int{1, 2, 2, 1, 1, 0})
	}
}
