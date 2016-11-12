package LeetCode

func InvertTree(root *TreeNode) *TreeNode{
	if root == nil{
		return nil
	}else{
		temp := root.Left
		root.Left = InvertTree(root.Right)
		root.Right = InvertTree(temp)
		return root
	}
}
