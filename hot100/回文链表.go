package main

func endOfFirstHalf(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	firstHalfEnd := endOfFirstHalf(head)
	secondHalfStart := reverseList(firstHalfEnd.Next)

	p1 := head
	p2 := secondHalfStart
	result := true
	for p2 != nil {
		if p1.Val != p2.Val {
			result = false
			break
		}
		p1 = p1.Next
		p2 = p2.Next
	}

	firstHalfEnd.Next = reverseList(secondHalfStart)
	return result
}