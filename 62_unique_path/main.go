package _2_unique_path

// TimeComplexity: O(n * m)
// SpaceComplexity: O(n * m)
func uniquePaths(m int, n int) int {
	// dp[row][col] returns number of way to get that coordinate.
	// if dp[1][1] returns 2, it means there are tow ways to get row = 1, col = 1.
	dp := make([][]int, m)
	for i, _ := range dp {
		dp[i] = make([]int, n)
	}

	// O(n*m)
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if row == 0 || col == 0 {
				dp[row][col] = 1
				continue
			}
			dp[row][col] = dp[row-1][col] + dp[row][col-1]
		}
	}
	return dp[m-1][n-1]
}

// TimeComplexity: O(n * m)
// SpaceComplexity: O(n)
func uniquePaths_space_optimized(m int, n int) int {
	dp := make([]int, n)
	for i, _ := range dp {
		dp[i] = 1
	}
	for i := 1; i < m; i++ {
		newDP := make([]int, n)
		newDP[0] = 1
		for j := 1; j < n; j++ {
			newDP[j] = dp[j] + newDP[j-1]
		}
		dp = newDP
	}
	return dp[n-1]
}
