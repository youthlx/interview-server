package tree

func IsSymmetric(root *TreeNode) bool {
	return isSymmetric(root, root)
}

func isSymmetric(root1, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}
	if root1.Val != root2.Val {
		return false
	}
	return isSymmetric(root1.Left, root2.Right) && isSymmetric(root1.Right, root2.Left)
}

func IsSymmetricByQueue(root *TreeNode) bool {
	u, v := root, root
	q := []*TreeNode{u, v}
	for len(q) > 0 {
		u, v = q[0], q[1]
		q = q[2:]
		if u == nil && v == nil {
			continue
		}
		if u == nil || v == nil {
			return false
		}
		if u.Val != v.Val {
			return false
		}
		q = append(q, u.Left, v.Right, u.Right, v.Left)
	}
	return true
}
