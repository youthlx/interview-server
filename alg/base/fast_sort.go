package base

// haha
func Partition(nums []int, start, end int) int {
	pivot := nums[start]
	l, r := start, end
	for l<r {
		for l<r && nums[r] >= pivot {
			r--
		}
		nums[l] = nums[r]
		for l<r && nums[l] <= pivot {
			l++
		}
		nums[r] = nums[l]
	}
	nums[l] = pivot
	return l
}



// rebase test 1