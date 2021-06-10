// 518. 零钱兑换 II，中等，https://leetcode-cn.com/problems/coin-change-2/

package coin_change_2

func Change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1

	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}
	return dp[amount]
}
