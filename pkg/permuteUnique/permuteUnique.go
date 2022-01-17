package main

import (
	"fmt"
	"sort"
)

var res [][]int

func permuteUnique(nums []int) [][]int {
	res = [][]int{}

	sort.Ints(nums)

	backTrack(nums, 0, []int{}, map[int]bool{})
	return res
}

func backTrack(nums []int, index int, tmp []int, selected map[int]bool) {

	if len(nums) == index {
		res = append(res, append([]int{}, tmp...))
		return
	}

	for i, n := range nums {

		// 如果上一个已经选择完并取消选择了，而自己和上一个一样，那么就不选自己了。选了的话反而会重复一遍
		if selected[i] || (i > 0 && n == nums[i-1] && !selected[i-1]) {
			continue
		}

		tmp2 := append(tmp, n)

		selected[i] = true

		backTrack(nums, index+1, tmp2, selected)

		selected[i] = false

		// tmp = tmp[:len(tmp)-1]
	}
}

func main() {
	//fmt.Println(permuteUnique([]int{1, 1, 2}))
	fmt.Println(permuteUnique([]int{-1, 2, -1, 2, 1, -1, 2, 1}))
}
