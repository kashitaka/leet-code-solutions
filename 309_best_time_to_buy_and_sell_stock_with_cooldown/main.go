package _09_best_time_to_buy_and_sell_stock_with_cooldown

// there are 3 states:
// - P: positioning stocks,
// - NP: no positioning stocks, but ready to buy, and
// - C: cooldown.
// To reach each state,
// - P: hold stocks that already bought, or buy stock from no-posi state
// - NP: stay no-posi or finish cooldown time.
// - C: sell stocks from posi state, and make profits
//
// Therefore, the recurrences of i-th day for each state are
//
//	P_i = max(P_{i-1}, NP_{i-1} - prices[i]),
//	NP_i = max(NP_{i-1}, C_{i-1}),
//	C_i = P_{i-1} + prices[i]).
//
// Note that initializing P_0 = - prices[0] is required,
// since we have to buy stock to be the posi state in the first day.
//
// What if the cooldown time is variable `t`?
// The recurrences of np changes as:
// NP_i = max(NP_{i-1}, P_{i-t}).
//
// Running time is O(n), where n is the length of given prices.
// Space complexity is O(3n) = O(n).
func maxProfit(prices []int) int {
	n := len(prices)

	p := make([]int, n)
	np := make([]int, n)
	c := make([]int, n)

	p[0] = -prices[0]

	for i := 1; i < n; i++ {
		p[i] = max(p[i-1], np[i-1]-prices[i])
		np[i] = max(np[i-1], c[i-1])
		c[i] = p[i-1] + prices[i]
	}
	return max(p[n-1], np[n-1], c[n-1])
}
