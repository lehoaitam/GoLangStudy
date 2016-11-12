package LeetCode



func MaxDepth(root *TreeNode) int{
	if root == nil{
		return 0
	}else{
		return MaxDepthBackTrack(root,1)
	}
}
func MaxDepthBackTrack(node *TreeNode, i int) int{
	if node.Left == nil && node.Right == nil{
		return i
	}else{
		maxLeft, maxRight := 0,0
		if node.Left != nil{
			maxLeft = MaxDepthBackTrack(node.Left, i + 1)
		}
		if node.Right != nil{
			maxRight = MaxDepthBackTrack(node.Right, i + 1)
		}
		if maxLeft > maxRight{
			return maxLeft
		}else {
			return maxRight
		}
	}
}