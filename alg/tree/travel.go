package tree

func PreorderByStack(root *TreeNode) []int {
	var stack []*TreeNode
	var res []int
	if root == nil {
		return res
	}
	stack = append(stack, root)
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		res = append(res, top.Val)
		stack = stack[:len(stack)-1] // handle
		if top.Right != nil {
			stack = append(stack, top.Right)
		}
		if top.Left != nil {
			stack = append(stack, top.Left)
		}
	}
	return res
}

func InorderByStack(root *TreeNode) []int {
	var stack []*TreeNode
	var res []int
	if root == nil {
		return res
	}
	cur := root
	for cur != nil || len(stack) > 0 {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, cur.Val) //handle
			cur = cur.Right
		}
	}
	return res
}

func PostorderByStack(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	var stack []*TreeNode
	stack = append(stack, root)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		res = append(res, node.Val)
		stack = stack[:len(stack)-1]
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	reverseArray(res)
	return res
}

func reverseArray(nums []int) {
	n := len(nums)
	for i := 0; i < n/2; i++ {
		nums[i], nums[n-1-i] = nums[n-1-i], nums[i]
	}
}
