package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{0, 0, 0}
	fmt.Println(threeSum(nums))

	nums = []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}

func threeSum(nums []int) [][]int {

	n := len(nums)

	if n < 3 {
		return [][]int{}
	}

	sort.Ints(nums)

	res := [][]int{}

	last := nums[0] - 1
	for i := 0; i < n-2; i++ {
		num := nums[i]
		if num == last {
			continue
		}
		last = num

		left := 0 - num

		lo, hi := i+1, n-1

		for lo < hi {
			nlo, nhi := nums[lo], nums[hi]
			if nlo+nhi == left {
				res = append(res, []int{num, nlo, nhi})
				hi -= 1
				for lo < hi && nums[hi] == nhi {
					hi -= 1
				}
				lo += 1
				for lo < hi && nums[lo] == nlo {
					lo += 1
				}
			} else if nlo+nhi > left {
				hi -= 1
				for lo < hi && nums[hi] == nhi {
					hi -= 1
				}
			} else {
				lo += 1
				for lo < hi && nums[lo] == nlo {
					lo += 1
				}
			}

		}
	}

	return res
}
