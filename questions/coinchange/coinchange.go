package coinchange

// You are given an integer array coins representing coins of different denominations
// and an integer amount representing a total amount of money.
// Return the fewest number of coins that you need to make up that amount. If that amount of
// money cannot be made up by any combination of the coins, return -1.
// You may assume that you have an infinite number of each kind of coin.

// Thoughts
// 1. This is another DP problem where want to cache the minimum number of coins to produce values
// 0...n - 1. At n, we iterate through all of the coins and find the coin that will give us the minimum
// number of coins + 1 given table[target - coinValue].
func Algo(coins []int, amount int) int {
	table := make([]int, amount+1)
	for curr := 1; curr <= amount; curr++ {
		minNumberCoins := 0
		minSet := false
		for _, coin := range coins {
			cachedValue := curr - coin
			if cachedValue >= 0 && table[cachedValue] != -1 {
				if minSet && table[cachedValue] < minNumberCoins {
					minNumberCoins = table[cachedValue]
				} else if !minSet {
					minNumberCoins = table[cachedValue]
					minSet = true
				}
			}
		}

		if minSet {
			table[curr] = minNumberCoins + 1
		} else {
			table[curr] = -1
		}
	}
	return table[amount]
}
