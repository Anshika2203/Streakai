# DFS API


Develop an API that accepts a set of directed graph edges, along with a start node and an end node, and returns all possible paths from the start node to the end node using Depth First Search (DFS).

### API Specification
Endpoint: /find-paths

Method: POST

Request Body:

The API accepts a JSON object with the following structure:
```
{
  "edges": [[0, 1], [0, 2], [1, 2], [1, 3], [2, 3], [3, 4]],
  "start": 0,
  "end": 4
}
```

edges: An array of edges representing the directed graph. Each edge is a two-element array [u, v] indicating a directed edge from node u to node v.

start: The starting node of the path.

end: The target node of the path.

## To Run the code

1. `go run main.go`

2. sent POST request to https://localhost:8080 with above json body.

