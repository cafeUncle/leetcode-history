package main

/**
全排列
https://leetcode-cn.com/problems/permutations/
*/

import "fmt"

var res = [][]int{}

func permute(nums []int) [][]int {
	res = [][]int{}

	backTrack(nums, 0, []int{}, map[int]bool{})

	return res
}

func backTrack(nums []int, index int, tmp []int, selected map[int]bool) {
	if len(nums) == index {
		res = append(res, tmp)
		return
	}

	for _, n := range nums {
		if selected[n] {
			continue
		}
		tmp2 := append(tmp, n)
		selected[n] = true
		backTrack(nums, index+1, tmp2, selected)
		selected[n] = false
		// 如果 append tmp 后没有初始一个 tmp2 传入 backTrack，而是直接把append结果赋给tmp，
		// 那么 tmp 的值就会修改，这里就需要进行 取消选择 了
		//tmp = tmp[:len(tmp)-1]
	}
}

func main() {
	fmt.Println(permute([]int{5, 4, 6, 2}))
}
