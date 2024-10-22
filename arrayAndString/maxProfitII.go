package arrayAndString

func maxProfitII(prices []int) int {
	buyPrice := prices[0]
	summa := 0

	for i := 1; i < len(prices); i++ {
		if prices[i-1] >= prices[i] {
			if prices[i-1] > buyPrice {
				summa += prices[i-1] - buyPrice
			}
			buyPrice = prices[i]

		}
	}
	return summa + prices[len(prices)-1] - buyPrice
}
