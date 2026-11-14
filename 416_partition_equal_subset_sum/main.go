package _16_partition_equal_subset_sum

// Let n be the size `nums` and t be the target number
// that can be obtained by t = (sum of nums)/2
// The running time of DP solution is O(nt).
// There several way to make the algorithm with O(nt) time,
// but the same O(nt) algorithms not necessarily give execution time.

// Reverse order DP with early return
// This is the MOST OPTIMAL solution with average 7 msec running time.
// The time complexity is O(nt)
func canPartition(nums []int) bool {
	total := 0
	for _, num := range nums {
		total += num
	}
	if total%2 != 0 {
		return false
	}

	target := total / 2

	dp := make([]bool, target+1)
	dp[0] = true
	for _, num := range nums {
		for i := target; i >= num; i-- {
			if dp[i-num] {
				dp[i] = true
			}
		}
		if dp[target] {
			return true
		}
	}
	return dp[target]
}

// The reverse order DP is somewhat difficult to understand,
// and the regular order DP may be easier, which is still O(nt) running time,
// but it requires copy of DP list in the beginning of each iteration,
// so the average time is 9 msec, slightly slower than  reverse DP.
func canPartition_asc_with_early_return(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%2 != 0 {
		return false
	}

	target := sum / 2

	prev := make([]bool, target+1)
	next := make([]bool, target+1)
	prev[0] = true

	for _, num := range nums {
		copy(next, prev)

		for s := num; s <= target; s++ {
			if prev[s-num] {
				next[s] = true
			}
		}
		if next[target] {
			return true
		}

		prev, next = next, prev
	}

	return prev[target]
}

// This might be the most straightforward solution with regular order.
// Although the running time is still O(nt), it takes 25 msec,
// which is slower than solutions above.
func canPartition_asc(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%2 != 0 {
		return false
	}

	target := sum / 2

	prev := make([]bool, target+1)
	prev[0] = true

	for _, num := range nums {
		next := make([]bool, target+1)
		for s := 0; s <= target; s++ {
			next[s] = prev[s]
			if s >= num && prev[s-num] {
				next[s] = true
			}
		}
		prev = next
	}

	return prev[target]
}

// We can make DP with hashmap, and it's still O(nt) running time,
// but it takes 85 msec, which implies a huge overhead for dealing
// with a hashmap.
func canPartition_hashmap(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%2 == 1 {
		return false
	}
	target := sum / 2

	hash := make(map[int]bool)
	hash[0] = true

	for _, num := range nums {
		keys := make([]int, 0, len(hash))
		for k, _ := range hash {
			keys = append(keys, k)
		}
		for _, v := range keys {
			hash[v+num] = true
			if v+num == target {
				return true
			}
		}
	}
	return false
}
