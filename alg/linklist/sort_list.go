package linklist

func SortList(head *ListNode) *ListNode {
	return sortList(head, nil)
}

/**
 * 自顶向下归并，循环不变量->左闭右开
 * O(NlogN), O(logN)
 */
func sortList(head, tail *ListNode) *ListNode {
	// 截止条件，切分后不剩节点或剩一个节点默认有序
	if head == nil || head.Next == nil {
		return head
	}
	// 通过快慢指针去寻找中间为止，慢指针一次一步，快指针一次两步
	// ----- 寻找中点模式 -----
	slow, fast := head, head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}
	}
	// ----- 寻找中点模式 -----
	mid := slow
	// 递归左子链，head到mid.prev
	lList := sortList(head, mid)
	// 递归右子链，mid到tail.prev
	rList := sortList(mid, tail)
	// 合并
	return merge(lList, rList)
}

func merge(head1, head2 *ListNode) *ListNode {
	// --- 合并两个有序链表模式 ---
	dummyHead := &ListNode{}
	p := head1
	q := head2
	dh := dummyHead
	for p != nil && q != nil {
		if p.Val < q.Val {
			dh.Next = p
			p = p.Next
		} else {
			dh.Next = q
			q = q.Next
		}
		dh = dh.Next
	}
	if p != nil {
		dh.Next = p
	}
	if q != nil {
		dh.Next = q
	}
	return dh.Next
	// --- 合并两个有序链表模式 ---
}

/**
 * 自底向上的归并排序，左闭右闭
 */
func SortList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var length int
	for node := head; node != nil; node = node.Next {
		length++
	}
	dummyHead := &ListNode{Next: head}
	for subLen := 1; subLen < length; subLen <<= 1 {
		prev, cur := dummyHead, dummyHead.Next
		for cur != nil {
			head1 := cur
			for i := 1; i < subLen && cur.Next != nil; i++ {
				cur = cur.Next
			}
			head2 := cur.Next
			cur.Next = nil
			cur = head2
			for i := 1; i < subLen && cur != nil && cur.Next != nil; i++ {
				cur = cur.Next
			}
			var next *ListNode
			if cur.Next != nil {
				next = cur.Next
				cur.Next = nil
			}
			prev.Next = merge(head1, head2)
			for prev.Next != nil {
				prev = prev.Next
			}
			cur = next
		}
	}
	return dummyHead.Next
}
