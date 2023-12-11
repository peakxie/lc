package main

import (
	"fmt"
	"math"
)

func main() {
	// 定义一个图
	graph := map[int][]Edge{
		0: {{1, 1}, {3, 3}},
		1: {{2, 2}, {3, 2}},
		2: {{5, 1}},
		3: {{4, 2}},
		4: {{5, 3}},
		5: {},
	}

	// 计算从节点0到其他节点的最短距离
	distances := dijkstra(graph, 0)
	fmt.Println("最短距离:", distances) // 输出: 最短距离: map[0:0 1:1 2:3 3:3 4:5 5:4]
}

type Edge struct {
	to     int
	weight int
}

func dijkstra(graph map[int][]Edge, start int) map[int]int {

	dist := make(map[int]int)
	unprocess := make(map[int]int)
	for k, _ := range graph {
		dist[k] = math.MaxInt32
		unprocess[k] = 0
	}
	dist[start] = 0

	for len(unprocess) > 0 {
		// 找到最小的路径
		minDist := math.MaxInt32
		minNode := -1
		for k, _ := range unprocess {
			if dist[k] < minDist {
				minDist = dist[k]
				minNode = k
			}
		}

		for _, v := range graph[minNode] {
			if _, exist := unprocess[v.to]; exist {
				newDist := dist[minNode] + v.weight
				if dist[v.to] > newDist {
					dist[v.to] = newDist
				}
			}
		}

		delete(unprocess, minNode)

	}
	return dist
}
