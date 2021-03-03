package main

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	temp := new(ListNode)
	temp.Next = head

	p := temp
	q := temp.Next
	for p != q {
		if q == nil || q.Next == nil {
			return false
		}
		p = p.Next
		q = q.Next.Next
	}
	return true
}
