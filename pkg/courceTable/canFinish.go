package main

import "fmt"

/**
课程表
https://leetcode-cn.com/problems/course-schedule/
*/
func main() {
	fmt.Println(canFinish(2, [][]int{{1, 0}}))
}

var onPath = make(map[int]bool)
var visited = make(map[int]bool)
var hasCycle = false

func canFinish(numCourses int, prerequisites [][]int) bool {
	onPath = make(map[int]bool)
	visited = make(map[int]bool)
	graph := make([][]int, numCourses)
	hasCycle = false

	for _, p := range prerequisites {
		graph[p[1]] = append(graph[p[1]], p[0])
	}

	for i, p := range graph {
		fmt.Println(i, " 被依赖于 ", p)
	}
	// 邻接表graph，中的每行，i 表示课程，p表示课程 i 被哪几个课程依赖。被依赖课程的被依赖课程中包含该课程时，认为有环。

	for course := range graph {
		// 判断每个课 course 是否可以完成学习
		traverse(graph, course)
	}
	return !hasCycle
}

func traverse(graph [][]int, courseName int) {
	if onPath[courseName] {
		// 发现环！！！  表示查询被依赖路径过程中，有入栈未出栈的内容发生了重复
		hasCycle = true
	}
	// visited[courseName] is true 表示这行已经扫描了，只需要扫描其它行
	if visited[courseName] {
		return
	}
	// 将节点 s 标记为已遍历
	visited[courseName] = true
	// 开始遍历节点 courseName
	onPath[courseName] = true
	for _, t := range graph[courseName] {
		// 去扫描依赖 courseName 的其他课
		traverse(graph, t)
	}
	// 节点 courseName 遍历完成
	onPath[courseName] = false
}
