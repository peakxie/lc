package top150

import "math"

func ladderLength(beginWord string, endWord string, wordList []string) int {

	graph := [][]int{}
	wordMap := map[string]int{}

	var addNode func(string) int
	addNode = func(s string) int {
		v, ok := wordMap[s]
		if !ok {
			v = len(wordMap)
			wordMap[s] = v
			graph = append(graph, []int{})
		}
		return v
	}

	var addEdge func(string) int
	addEdge = func(s string) int {
		id := addNode(s)
		bs := []byte(s)
		for i, v := range bs {
			bs[i] = '*'
			id1 := addNode(string(bs))
			graph[id] = append(graph[id], id1)
			graph[id1] = append(graph[id1], id)
			bs[i] = v
		}
		return id
	}

	for _, v := range wordList {
		addEdge(v)
	}

	id := addEdge(beginWord)
	id1, ok := wordMap[endWord]
	if !ok {
		return 0
	}

	countList := make([]int, len(wordMap))
	for i, _ := range countList {
		countList[i] = math.MaxInt
	}

	countList[id] = 0
	queue := []int{id}
	for len(queue) > 0 {
		id := queue[0]
		queue = queue[1:]

		for _, v := range graph[id] {
			if countList[v] == math.MaxInt {
				countList[v] = countList[id] + 1
				queue = append(queue, v)
			}
		}
	}
	if countList[id1] == math.MaxInt {
		return 0
	}
	return countList[id1]/2 + 1
}
