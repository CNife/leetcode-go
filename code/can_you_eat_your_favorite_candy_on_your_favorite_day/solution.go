// 1744. 你能在你最喜欢的那天吃到你最喜欢的糖果吗？，中等
// https://leetcode-cn.com/problems/can-you-eat-your-favorite-candy-on-your-favorite-day/

package can_you_eat_your_favorite_candy_on_your_favorite_day

func CanEat(candiesCount []int, queries [][]int) []bool {
	result := make([]bool, len(queries))

	candies := make([]int64, len(candiesCount))
	candies[0] = int64(candiesCount[0])
	for i := 1; i < len(candies); i++ {
		candies[i] = candies[i-1] + int64(candiesCount[i])
	}

	for i, query := range queries {
		candyType, day, dayMaxEatNum := query[0], int64(query[1]+1), int64(query[2])
		result[i] = (candyType == 0 && day <= candies[candyType]) ||
			(candyType != 0 && day <= candies[candyType] && day*dayMaxEatNum > candies[candyType-1])
	}
	return result
}
