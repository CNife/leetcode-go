package dungeon_game

func CalculateMinimumHP(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])
	memo := make([][]int, 0, m)
	for i := 0; i < m; i++ {
		memo = append(memo, make([]int, n))
	}

	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i == m-1 && j == n-1 {
			return maxOf1(1 - dungeon[i][j])
		}
		if memo[i][j] > 0 {
			return memo[i][j]
		}

		var result int
		switch {
		case i == m-1:
			result = maxOf1(dfs(i, j+1) - dungeon[i][j])
		case j == n-1:
			result = maxOf1(dfs(i+1, j) - dungeon[i][j])
		default:
			result = maxOf1(min(dfs(i+1, j), dfs(i, j+1)) - dungeon[i][j])
		}
		memo[i][j] = result
		return result
	}
	return dfs(0, 0)
}

func maxOf1(num int) int {
	if num < 1 {
		return 1
	}
	return num
}

func min(lhs, rhs int) int {
	if lhs < rhs {
		return lhs
	}
	return rhs
}
