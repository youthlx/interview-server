package array

func MaxProfit(prices []int) int {
	minPrice := 10001
	maxDiff := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
		diff := prices[i] - minPrice
		if diff > maxDiff {
			maxDiff = diff
		}
	}
	return maxDiff
}
