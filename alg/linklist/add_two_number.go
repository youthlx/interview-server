package linklist

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	p := l1
	q := l2
	h := dummyHead
	var add int
	for p != nil && q != nil {
		sum := (p.Val + q.Val + add) % 10
		add = (p.Val + q.Val + add) / 10
		h.Next = &ListNode{Val: sum}
		p = p.Next
		q = q.Next
		h = h.Next
	}
	for p != nil {
		sum := (p.Val + add) % 10
		add = (p.Val + add) / 10
		h.Next = &ListNode{Val: sum}
		p = p.Next
		h = h.Next
	}
	for q != nil {
		sum := (q.Val + add) % 10
		add = (q.Val + add) / 10
		h.Next = &ListNode{Val: sum}
		q = q.Next
		h = h.Next
	}
	if add == 1 {
		h.Next = &ListNode{Val: 1}
	}
	return dummyHead.Next
}
