package main

/**
剑指 Offer 12. 矩阵中的路径
https://leetcode-cn.com/problems/ju-zhen-zhong-de-lu-jing-lcof/
*/

import "fmt"

func main() {
	fmt.Println(exist([][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}, "ABCCED"))
}

var res = false

func exist(board [][]byte, word string) bool {
	res = false

	backTrack(board, 0, word, -1, -1)
	return res
}

func backTrack(board [][]byte, index int, word string, row int, col int) {

	if res {
		return
	}

	if index == len(word) {
		res = true
		return
	}

	for i, line := range board {

		if row != -1 && i != row && i != row+1 && i != row-1 {
			continue
		}

		for j, cell := range line {

			if board[i][j] == '0' {
				continue
			}

			if col != -1 && j != col && j != col+1 && j != col-1 {
				continue
			}

			if row != -1 && col != -1 && i != row && j != col {
				continue
			}

			if word[index] == cell {

				tmp := board[i][j]
				board[i][j] = '0'
				backTrack(board, index+1, word, i, j)
				board[i][j] = tmp
			}
		}

	}

}
