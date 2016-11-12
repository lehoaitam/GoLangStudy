package LeetCode

func SwapPairs(head *ListNode) *ListNode{
	if head == nil || head. Next == nil {
		return head
	}
	p1 := head
	p2 := p1.Next
	p3 := p2.Next

	p2.Next = p1
	p1.Next = SwapPairs(p3)
	return p2
}
