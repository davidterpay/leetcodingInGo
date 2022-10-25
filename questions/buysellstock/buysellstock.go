package buysellstock

// You are given an array prices where prices[i] is the price of a given stock on the ith day.
// You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.
// Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

// Thoughts
// 1. The iterative and noob solution would be to iterate day + 1, ..., n for each stock to see
// which of the days and end dates give you highest profit. This is noob because this takes O(n^2)
// time to compute.
// 2. At any given day, the way that we maximize profits is by finding some day before that that had
// the lowest price. Therefore, at any point we want to keep track of the smallest price and calcuate
// profit for that day and maximize for profit at every step while minimizing the cheapest price.
func Algo(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	minPrice := prices[0]
	profit := 0

	for _, price := range prices[1:] {
		if price < minPrice {
			minPrice = price
		} else {
			tempProfit := price - minPrice
			if tempProfit > profit {
				profit = tempProfit
			}
		}
	}

	return profit
}
