package string

//你有一个单词列表words和一个模式pattern，你想知道 words 中的哪些单词与模式匹配。
//
//如果存在字母的排列 p，使得将模式中的每个字母 x 替换为 p(x) 之后，我们就得到了所需的单词，那么单词与模式是匹配的。
//
//（回想一下，字母的排列是从字母到字母的双射：每个字母映射到另一个字母，没有两个字母映射到同一个字母。）
//
//返回 words 中与给定模式匹配的单词列表。
//
//你可以按任何顺序返回答案。
//
//示例：
//
//输入：words = ["abc","deq","mee","aqq","dkd","ccc"], pattern = "abb"
//输出：["mee","aqq"]
//解释：
//"mee" 与模式匹配，因为存在排列 {a -> m, b -> e, ...}。
//"ccc" 与模式不匹配，因为 {a -> c, b -> c, ...} 不是排列。
//因为 a 和 b 映射到同一个字母。
func FindAndReplacePattern(words []string, pattern string) []string {
	var ans []string
	for _, word := range words {
		// 一一对应
		if match(word, pattern) && match(pattern, word) {
			ans = append(ans, word)
		}
	}
	return ans
}

// match只表示word中的每个字符唯一的对应pattern中的一个
// 但是会存在以下场景 w：abb， p：ccc，这样的话a -> c, b->c
func match(word, pattern string) bool {
	m := make(map[byte]byte)
	for i := 0; i < len(word); i++ {
		if m[word[i]] == 0 {
			m[word[i]] = pattern[i]
			continue
		}
		// ---在此之前已经具备了一定的模式，且模式具备后效性---
		if m[word[i]] != pattern[i] {
			return false
		}
	}
	return true
}

//后效性：后效性是在遍历列表的场景下，如果遍历到位置i时，存储或者处理的结果对其后的处理或者结果仍然具备一定的效果
//相关的题目有：两数之和的map版本，股票I问题
