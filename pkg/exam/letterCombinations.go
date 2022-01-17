package main

import "fmt"

var phoneMap = map[byte][]string{
	'2': []string{"a", "b", "c"},
	'3': []string{"d", "e", "f"},
	'4': []string{"g", "h", "i"},
	'5': []string{"j", "k", "l"},
	'6': []string{"m", "n", "o"},
	'7': []string{"p", "q", "r", "s"},
	'8': []string{"t", "u", "v"},
	'9': []string{"w", "x", "y", "z"},
}

var res = []string{}

func letterCombinations(digits string) []string {
	res = []string{}

	if len(digits) == 0 {
		return res
	}

	backTrace(digits, 0, "")
	return res
}

func backTrace(digits string, index int, str string) {
	fmt.Println(len(digits), index)
	if len(digits) == index {
		res = append(res, str)
		fmt.Println("")
		return
	}

	options := phoneMap[digits[index]]

	for _, o := range options {
		backTrace(digits, index+1, str+o)
	}

}

func main() {
	fmt.Println(letterCombinations("23"))
}
