package LeetCode

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode{
	if root == nil{
		return nil
	}
	if root == p || root == q{
		return root
	}
	leftAncestor := lowestCommonAncestor(root.Left, p, q)
	rightAncestor := lowestCommonAncestor(root.Right, p, q)
	if leftAncestor != nil && rightAncestor != nil{
		return root
	}
	if leftAncestor != nil{
		return leftAncestor
	}else{
		return rightAncestor
	}


}
