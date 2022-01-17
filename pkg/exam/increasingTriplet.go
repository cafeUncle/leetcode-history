package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{1, 0, 0, 0, 0, 0, 100000000}
	fmt.Println(increasingTriplet(nums))
}

func increasingTriplet(nums []int) bool {
	lens := len(nums)

	if lens < 3 {
		return false
	}

	first := nums[0]
	second := int(math.Pow(2, 31) - 1)

	for i := 1; i < lens; i++ {
		fmt.Printf("in: f %d s %d n %d \n", first, second, nums[i])
		if nums[i] > second {
			return true
		} else if nums[i] < first {
			first = nums[i]
		} else if nums[i] > first && nums[i] < second {
			second = nums[i]
		}
		fmt.Printf("ou: f %d s %d n %d \n", first, second, nums[i])
	}

	return false
}
