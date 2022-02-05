package main

/**
和为 K 的最少斐波那契数字数目
https://leetcode-cn.com/problems/find-the-minimum-number-of-fibonacci-numbers-whose-sum-is-k/
*/

import "fmt"

func main() {
	fmt.Println(findMinFibonacciNumbers(7))
}
func findMinFibonacciNumbers(k int) int {
	a, b := 1, 1
	sorts := []int{a, b}
	for a+b < k {
		next := a + b
		sorts = append(sorts, next)
		a = b
		b = next
	}
	fmt.Println(sorts)
	count := 0
	idx := len(sorts) - 1
	for k > 0 {
		num := sorts[idx]
		idx--
		if k < num {
			continue
		}
		count++
		k -= num
	}
	return count
}
