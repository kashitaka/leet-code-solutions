package _139_word_break

// For a given s with length n,
// let dp be an array with size n + 1
// dp[i] is the solution of subproblem s[0:i]
// dp[0] is true, since s[0:0] = "" is valid.
// dp[n+1] contains the solution of the problem,
// as s[0:n+1] = s itself.
// Assume length of wordDict is m
// the running time is O(nm), in n iterations, check all m words.
func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 0; i < len(s); i++ {
		for _, word := range wordDict {
			l := len(word)
			if i+1 < l {
				continue
			}
			if s[i+1-l:i+1] == word && !dp[i+1] { // avoid to be overidden true value if it has.
				dp[i+1] = dp[i+1-l]
			}
		}
	}
	return dp[len(dp)-1]
}
