package LeetCode

func IsSameTree(p *TreeNode, q *TreeNode) bool{
	if p == nil && q == nil{
		return true
	}else if (p == nil && q!= nil ) || (p != nil && q == nil){
		return false
	}else if p.Val != q.Val{
		return false
	}else{
		return IsSameTree(p.Left,q.Left) && IsSameTree(p.Right, q.Right)
	}
}
