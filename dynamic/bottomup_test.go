package dynamic

import (
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestSolveSmallInput(t *testing.T) {
	coins := []int{1, 2, 3}
	res := solve(6, 3, coins)

	assert.Equal(t, 7, res)
}

func TestSolveLargeInput(t *testing.T) {
	coins := []int{16,30,9,17,40,13,42,5,25,49,7,23,1,44,4,11,33,12,27,2,38,24,28,32,14,50}
	res := solve(245, 26, coins)

	assert.Equal(t, 64027917156, res)
}
