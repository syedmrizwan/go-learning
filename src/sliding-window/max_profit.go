package sliding_window

func maxProfit(prices []int) int {
	max, left, right := 0, 0, 1

	for right < len(prices) {
		if prices[left] < prices[right] {
			profit := prices[right] - prices[left]
			if profit > max {
				max = profit
			}
		} else {
			left = right
		}

		right++
	}
	return max
}
