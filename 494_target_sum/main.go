package _94_target_sum

// Slice based solution shows BETTER performance with a LeetCode's test set.
// Running Time: O(2000 * n) = O(n), where n is the size of nums
// Space Complexity: at most O(2000) = O(1)
func findTargetSumWays(nums []int, target int) int {
	dp := make([]int, 2001)
	dp[1000] = 1

	for _, num := range nums {
		newDP := make([]int, 2001)
		for i, v := range dp {
			if v > 0 {
				newDP[i-num] += v
				newDP[i+num] += v
			}
		}
		dp = newDP
	}
	return dp[1000+target]
}


// Although the same running time, this is LESS EFFICIENT than the slice based solution,
// because accessing hash key requires computational workload to make its key.
// This may perform better with sparse input: small size array and huge numbers in it.
func findTargetSumWays_hash(nums []int, target int) int {
	dp := make(map[int]int)
	dp[0] = 1
	for i := 0; i < len(nums); i++ {
		newDP := make(map[int]int)
		num := nums[i]
		for k, v := range dp {
			newDP[k-num] += v
			newDP[k+num] += v
		}
		dp = newDP
	}
	return dp[target]
}
