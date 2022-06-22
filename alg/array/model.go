package array

// 滑动窗口的核心思想
// 滑动窗口算法的思路是这样：
// 1、我们在字符串 S 中使用双指针中的左右指针技巧，初始化 left = right = 0，把索引闭区间 [left, right] 称为一个「窗口」。
//
// 2、我们先不断地增加 right 指针扩大窗口 [left, right]，直到窗口中的字符串符合要求（包含了 T 中的所有字符）。
//
// 3、此时，我们停止增加 right，转而不断增加 left 指针缩小窗口 [left, right]，直到窗口中的字符串不再符合要求（不包含 T 中的所有字符了）。
//    同时，每次增加 left，我们都要更新一轮结果。
// int left = 0, right = 0;
//
// while (right < s.size()) {
//     window.add(s[right]);
//     right++;
//
//     while (valid) {
//         window.remove(s[left]);
//         left++;
//     }
//	}
