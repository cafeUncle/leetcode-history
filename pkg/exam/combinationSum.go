package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 3, 4, 5, 2, 6, 8}
	fmt.Println(combinationSum(nums, 7))
}

func combinationSum(candidates []int, target int) [][]int {
	n := len(candidates)
	sort.Ints(candidates)
	var res [][]int
	var tmp []int
	var dfs func(idx, remained int)
	dfs = func(idx, remained int) {
		if remained == 0 {
			res = append(res, append([]int{}, tmp...))
			return
		}
		if idx == n {
			return
		}
		if candidates[idx] > remained {
			return
		}
		dfs(idx+1, remained)
		tmp = append(tmp, candidates[idx])
		dfs(idx, remained-candidates[idx])
		tmp = tmp[:len(tmp)-1]
	}
	dfs(0, target)
	return res
}
