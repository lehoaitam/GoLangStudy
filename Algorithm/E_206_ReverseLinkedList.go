package LeetCode

func ReverseList(head *ListNode) *ListNode{
	if head == nil{
		return nil
	}else{
		var temp *ListNode
		temp = nil
		curNode := head
		nextNode := head
		for nextNode != nil{
			nextNode = nextNode.Next
			curNode.Next = temp
			temp = curNode
			curNode = nextNode
		}
		return temp

	}
}
