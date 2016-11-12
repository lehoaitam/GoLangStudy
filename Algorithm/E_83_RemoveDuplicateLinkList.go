package LeetCode
func DeleteDuplicates(head *ListNode) *ListNode {
	if head == nil{
		return nil
	}
	curNode := head
	nextNode := head
	for nextNode != nil{
		if nextNode != nil{
			if curNode.Val == nextNode.Val{
				nextNode = nextNode.Next
			}else{
				curNode.Next = nextNode
				curNode = nextNode
			}
		}
	}
	curNode.Next = nil
	return head
}
