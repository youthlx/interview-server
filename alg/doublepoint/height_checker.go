package doublepoint

//学校打算为全体学生拍一张年度纪念照。根据要求，学生需要按照 非递减 的高度顺序排成一行。
//
//排序后的高度情况用整数数组 expected 表示，其中 expected[i] 是预计排在这一行中第 i 位的学生的高度（下标从 0 开始）。
//
//给你一个整数数组 heights ，表示 当前学生站位 的高度情况。heights[i] 是这一行中第 i 位学生的高度（下标从 0 开始）。
//
//返回满足 heights[i] != expected[i] 的 下标数量 。
//输入：heights = [1,1,4,2,1,3]
//输出：3
//解释：
//高度：[1,1,4,2,1,3]
//预期：[1,1,1,2,3,4]
//下标 2 、4 、5 处的学生高度不匹配。
//
//输入：heights = [5,1,2,3,4]
//输出：5
//解释：
//高度：[5,1,2,3,4]
//预期：[1,2,3,4,5]
//所有下标的对应学生高度都不匹配。

func HeightChecker(heights []int) int {
	var n int
	for _, num := range heights {
		if num > n {
			n = num
		}
	}
	cs := make([]int, n+1)
	for _, num := range heights {
		cs[num]++
	}
	idx := 0
	var cnt int
	for i := 0; i < len(cs); i++ {
		times := cs[i]
		for times > 0 {
			// ---- 如果是要求出排序后的数组，那么将这段改为 ret = append(ret, i)
			if heights[idx] != i {
				cnt++
			}
			idx++
			// ----
			times--
		}
	}
	return cnt
}

// 技巧：计数排序，获取列表最大元素，假设为n，构建n+1的数组统计原数组的频率
// 场景：已知数组中最大元素和最小元素，且最大元素可接受
// 原数组: 1 1 4 2 1 3
// 频率:  0 3 1 1 1 （下标代表原数组元素，值为频率）
// 所以结果就是三个1，1个2，1个3，1个4然后用列表装起来就是
// 1 1 1 2 3 4
