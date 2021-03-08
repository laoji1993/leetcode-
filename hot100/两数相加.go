package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := new(ListNode)
	temp := res
	up := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + up
		sum, up = sum%10, sum/10
		temp.Next = &ListNode{Val: sum}
		temp = temp.Next
	}
	if up != 0 {
		temp.Next = &ListNode{Val: up}
	}
	return res.Next
}
