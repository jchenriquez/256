package main

import (
	"fmt"
	"math"
)

func key (i, j int) string {
	return fmt.Sprintf("%d,%d", i, j)
}

func minCostHelp(i, prevSelectedIndex int, costs [][]int, seen map[string]int)int {

	ky := key(i, prevSelectedIndex)

	if value, saw := seen[ky]; saw {
		return value
	}

	min := math.MaxInt32

	for j := 0; j < len(costs[i]); j++ {

		if j != prevSelectedIndex {
			if i < len(costs)-1 {
				min = int(math.Min(float64(min), float64(minCostHelp(i+1, j, costs, seen) + costs[i][j])))
			} else {
				min = int(math.Min(float64(min), float64(costs[i][j])))
			}
		}

	}

	seen[ky] = min
	return min
}

func minCost(costs [][]int) int {

	if len(costs) == 0 {
		return 0
	}

	seen := make(map[string] int, len(costs)*len(costs[0]))

	return minCostHelp(0, -1, costs, seen)

}

func main() {
	fmt.Printf("min cost is %d\n", minCost([][]int{{5, 8, 6}, {19, 14, 13}, {7, 5, 12}, {14, 15, 17}, {3, 20, 10}}))
}
