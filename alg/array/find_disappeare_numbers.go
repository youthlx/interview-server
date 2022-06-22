package array

func FindDisappearedNumbers(nums []int) []int {
	a := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		a[nums[i]]++
	}
	var ret []int
	for i := 1; i < len(a); i++ {
		if a[i] == 0 {
			ret = append(ret, i)
		}
	}
	return ret
}

func FindDisappearedNumbersFast(nums []int) []int {
	var ret []int
	n := len(nums)
	for _, v := range nums {
		v = (v - 1) % n
		nums[v] += n
	}
	for i, v := range nums {
		if v <= n {
			ret = append(ret, i+1)
		}
	}
	return ret
}
