func maxProfit(prices []int) int {
    buy, sell := 0, 1
	maxProfit := 0

	for sell < len(prices) {
		if prices[buy] < prices[sell] {
			profit := prices[sell] - prices[buy]
			if maxProfit < profit {
				maxProfit = profit
			}
		} else {
			buy = sell
		}
		sell++
	}

	return maxProfit
}
