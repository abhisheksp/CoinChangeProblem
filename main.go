package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	N, M, coins := input()
	res := solve(N, M, coins)
	fmt.Println(res)
}

// Copy from dynamic package
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

func input() (int, int, []int) {
	r := bufio.NewReader(os.Stdin)
	nStr, _ := r.ReadString('\n')

	nStr = strings.TrimSpace(nStr)
	line := strings.Split(nStr, " ")

	N, _ := strconv.Atoi(line[0])
	M, _ := strconv.Atoi(line[1])

	coinsLine, _ := r.ReadString('\n')
	trimmed_stream := strings.TrimSpace(coinsLine)
	coinsStr := strings.Split(trimmed_stream, " ")

	coins := []int{}
	for i := 0; i < M; i++ {
		coin, _ := strconv.Atoi(coinsStr[i])

		coins = append(coins, coin)
	}

	return N, M, coins
}
