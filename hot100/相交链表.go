package main

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	temp1,temp2 := headA,headB
	for temp1 != temp2 {
		if temp1 != nil {
			temp1 = temp1.Next
 		} else {
 			temp1 = headB
		}
		if temp2 != nil {
			temp2 = temp2.Next
		} else {
			temp2 = headA
		}
	}
	if temp1 != nil {
		return temp1
	}
	return nil
}
