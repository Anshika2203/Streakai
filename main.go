package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type RequestBody struct {
	Edges [][]int `json:"edges"`
	Start int     `json:"start"`
	End   int     `json:"end"`
}

// {
// 	"edges": [[0, 1], [0, 2], [1, 2], [1, 3], [2, 3], [3, 4]],
// 	"start": 0,
// 	"end": 4
//   }

type ResponseBody struct {
	Path [][]int `json:"path"`
}

// eg
// {
//   "paths": [
//     [0, 1, 3, 4],
//     [0, 2, 3, 4],
//     [0, 1, 2, 3, 4],
//   ]
// }

func buildGraph(edges [][]int) map[int][]int {
	graph := make(map[int][]int)
	for _, edges := range edges {
		u, v := edges[0], edges[1]
		graph[u] = append(graph[u], v)
	}
	return graph
}

func findPathsDFS(graph map[int][]int, current, end int, path []int, result *[][]int, visited map[int]bool) {
	if current == end {
		pathCopy := make([]int, len(path))
		copy(pathCopy, path)
		*result = append(*result, pathCopy)
		return
	}
	visited[current] = true
	for _, neighbors := range graph[current] {
		if !visited[neighbors] {
			findPathsDFS(graph, neighbors, end, append(path, neighbors), result, visited)
		}
	}
	visited[current] = false
}

func findPathHandler(w http.ResponseWriter, r *http.Request) {
	var reqBody RequestBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	graph := buildGraph(reqBody.Edges)
	var result [][]int
	visited := make(map[int]bool)
	findPathsDFS(graph, reqBody.Start, reqBody.End, []int{reqBody.Start}, &result, visited)
	response := ResponseBody{
		Path: result,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/find-paths", findPathHandler)
	fmt.Println("Server running on local8080")
	http.ListenAndServe(":8080", nil)
}
