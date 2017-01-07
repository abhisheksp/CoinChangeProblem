package dynamic

func solve(N int, M int, coins []int) int {
	result := make([][]int, 0)
	for i := 0; i <= M; i++ {
		result = append(result, make([]int, N+1))
	}

	result[0][0] = 1
	coins = append([]int{0}, coins...)
	for i := 1; i <= M; i++ {
		for j := 0; j <= N; j++ {
			if j < coins[i] {
				result[i][j] = result[i-1][j]
			} else {
				result[i][j] = result[i-1][j] + result[i][j-coins[i]]
			}
		}
	}
	return result[M][N]
}