package problem121

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	priceBuy := prices[0]
	priceProfit := 0

	for i := 1; i < len(prices); i++ {
		if priceBuy > prices[i] {
			priceBuy = prices[i]
		} else if prices[i]-priceBuy > priceProfit {
			priceProfit = prices[i] - priceBuy
		}
	}

	return priceProfit
}
