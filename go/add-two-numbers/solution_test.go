// BEGIN: 8f7e5d4a8c3b
package addtwonumbers

import (
	"reflect"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		l1   *ListNode
		l2   *ListNode
		want *ListNode
	}{
		{
			&ListNode{2, &ListNode{4, &ListNode{3, nil}}},
			&ListNode{5, &ListNode{6, &ListNode{4, nil}}},
			&ListNode{7, &ListNode{0, &ListNode{8, nil}}},
		},
		{
			&ListNode{0, nil},
			&ListNode{0, nil},
			&ListNode{0, nil},
		},
		{
			&ListNode{9, &ListNode{9, &ListNode{9, nil}}},
			&ListNode{9, &ListNode{9, &ListNode{9, nil}}},
			&ListNode{8, &ListNode{9, &ListNode{9, &ListNode{1, nil}}}},
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := addTwoNumbers(tt.l1, tt.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddTwoNumbersV2(t *testing.T) {
	tests := []struct {
		l1   *ListNode
		l2   *ListNode
		want *ListNode
	}{
		{
			&ListNode{2, &ListNode{4, &ListNode{3, nil}}},
			&ListNode{5, &ListNode{6, &ListNode{4, nil}}},
			&ListNode{7, &ListNode{0, &ListNode{8, nil}}},
		},
		{
			&ListNode{0, nil},
			&ListNode{0, nil},
			&ListNode{0, nil},
		},
		{
			&ListNode{9, &ListNode{9, &ListNode{9, nil}}},
			&ListNode{9, &ListNode{9, &ListNode{9, nil}}},
			&ListNode{8, &ListNode{9, &ListNode{9, &ListNode{1, nil}}}},
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := addTwoNumbersV2(tt.l1, tt.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbersV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

// END: 8f7e5d4a8c3b
