package linklist

// DetectCycleII O(N), O(1)
func DetectCycleII(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		if fast == slow {
			start := head
			for start != slow {
				start = start.Next
				slow = slow.Next
			}
			return slow
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return nil
}
