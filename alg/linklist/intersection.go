package linklist

func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	var len1, len2 int
	la := headA
	for la != nil {
		len1++
		la = la.Next
	}
	lb := headB
	for lb != nil {
		len2++
		lb = lb.Next
	}
	pa := headA
	pb := headB
	if len1 > len2 {
		for i := 0; i < len1-len2; i++ {
			pa = pa.Next
		}
	} else {
		for i := 0; i < len2-len1; i++ {
			pb = pb.Next
		}
	}
	for pa != nil && pb != nil {
		if pa == pb {
			return pa
		}
		pa = pa.Next
		pb = pb.Next
	}
	return nil
}
