package base

// Search()方法回使用“二分查找”算法来搜索某指定切片[0:n]，
// 并返回能够使f(i)=true的最小的i（0<=i<n）值，并且会假定，如果f(i)=true，则f(i+1)=true，即对于切片[0:n]，
// i之前的切片元素会使f()函数返回false，i及i之后的元素会使f()函数返回true。
// 但是，当在切片中无法找到时f(i)=true的i时（此时切片元素都不能使f()函数返回true），Search()方法会返回n（而不是返回-1)
func search(n int, f func(int) bool) int {
	i, j := 0, n
	for i < j {
		h := int(uint(i+j) >> 1)
		if !f(h) {
			i = h + 1
		} else {
			j = h
		}
	}
	return i
}
