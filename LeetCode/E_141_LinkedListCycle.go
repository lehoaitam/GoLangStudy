package LeetCode

func HasCycle(head *ListNode) bool{
	if head == nil{
		return false
	}
	slowCursor := head
	fastCursor := head.Next
	cnt := 1
	lim := 2
	for slowCursor != fastCursor{
		if fastCursor == nil || slowCursor == nil{
			return false
		}
		if cnt == lim{
			cnt = 1
			lim = lim * 2
			slowCursor = fastCursor
		}else{
			cnt++
		}
		fastCursor = fastCursor.Next
	}
}
