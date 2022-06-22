package tree

func InOrderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		res = append(res, root.Val)
		dfs(root.Right)
	}
	dfs(root)
	return res
}
