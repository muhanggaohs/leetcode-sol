func longestSubsequence(arr []int, difference int) int {
    dp := make(map[int]int)
    maxLength := 1

    for _, num := range arr {
        if dp[num-difference] > 0 {
            dp[num] = dp[num-difference] + 1
        } else {
            dp[num] = 1
        }
        maxLength = max(maxLength, dp[num])
    }

    return maxLength
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
