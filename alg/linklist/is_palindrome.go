package linklist

func IsPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	firstHalfEnd := findMid(head)
	secondHalfStart := reverse(firstHalfEnd.Next)
	p1 := head
	p2 := secondHalfStart
	result := true
	for p1 != nil && p2 != nil {
		if p1.Val != p2.Val {
			result = false
			break
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	firstHalfEnd.Next = reverse(secondHalfStart)
	return result
}

func findMid(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}
