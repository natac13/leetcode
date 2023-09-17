package addtwonumbers

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	cl1 := l1
	cl2 := l2

	var head *ListNode
	var current *ListNode
	var carry int

	for cl1 != nil || cl2 != nil {

		var sum int

		if cl1 != nil {
			sum += cl1.Val
			cl1 = cl1.Next
		}

		if cl2 != nil {
			sum += cl2.Val
			cl2 = cl2.Next
		}

		sum += carry

		carry = sum / 10
		sum = sum % 10

		if current == nil {
			current = &ListNode{Val: sum}
			head = current
		} else {
			current.Next = &ListNode{Val: sum}
			current = current.Next
		}

	}

	if carry > 0 {
		current.Next = &ListNode{Val: carry}
	}

	return head
}

func addTwoNumbersV2(l1, l2 *ListNode) *ListNode {
	// create a new linked list to store the result
	result := &ListNode{}
	// set the current node to the head of the result list
	current := result

	// loop through both input lists until both are nil
	for l1 != nil || l2 != nil {
		if l1 != nil {
			// add the value of the current node in l1 to the current node in the result list
			current.Val += l1.Val
			// move to the next node in l1
			l1 = l1.Next
		}

		if l2 != nil {
			// add the value of the current node in l2 to the current node in the result list
			current.Val += l2.Val
			// move to the next node in l2
			l2 = l2.Next
		}
		// if the sum of the current nodes is greater than or equal to 10
		if current.Val >= 10 {
			// subtract 10 from the current node in the result list
			current.Val -= 10
			// set the next node in the result list to a new node with value 1
			current.Next = &ListNode{Val: 1}
			// if the sum of the current nodes is less than 10 and there are more nodes in either l1 or l2
		} else if l1 != nil || l2 != nil {
			// set the next node in the result list to a new node with value 0
			current.Next = &ListNode{}
		}

		// move to the next node in the result list
		current = current.Next
	}
	return result // return the head of the result list
}
