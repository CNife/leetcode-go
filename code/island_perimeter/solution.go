package island_perimeter

func IslandPerimeter(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	rowCount, columnCount := make(chan int), make(chan int)

	go func() {
		rowResult := 0
		for i := 0; i < m; i++ {
			flag := 0
			for j := 0; j < n; j++ {
				if v := grid[i][j]; v != flag {
					flag = v
					rowResult++
				}
			}
			if flag == 1 {
				rowResult++
			}
		}
		rowCount <- rowResult
	}()

	go func() {
		columnResult := 0
		for j := 0; j < n; j++ {
			flag := 0
			for i := 0; i < m; i++ {
				if v := grid[i][j]; v != flag {
					flag = v
					columnResult++
				}
			}
			if flag == 1 {
				columnResult++
			}
		}
		columnCount <- columnResult
	}()

	return <-rowCount + <-columnCount
}
