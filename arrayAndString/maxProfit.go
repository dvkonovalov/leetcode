package arrayAndString

func maxProfit(prices []int) int {
	buyPrice := prices[0]
	maxProf := 0

	for i := 1; i < len(prices); i++ {
		if buyPrice > prices[i] {
			buyPrice = prices[i]
		}
		if maxProf < prices[i]-buyPrice {
			maxProf = prices[i] - buyPrice
		}
	}
	return maxProf
}
