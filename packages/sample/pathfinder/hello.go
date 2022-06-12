package main

type Request struct {
	Start []int   `json:"start"`
	End   []int   `json:"end"`
	Grid  [][]int `json:"grid"`
}

type Response struct {
	StatusCode int               `json:"statusCode,omitempty"`
	Headers    map[string]string `json:"headers,omitempty"`
	Body       Body              `json:"body,omitempty"`
}

type Body struct {
	Path [][]int `json:"path"`
}

type Position struct {
	X int
	Y int
}

type Node struct {
	parent   *Node
	position Position
	g        int
	h        int
	f        int
}

func Main(in Request) (*Response, error) {

	g := in.Grid
	start := Position{in.Start[0], in.Start[1]}
	end := Position{in.End[0], in.End[1]}

	return &Response{
		Body: Body{
			Path: FindPath(start, end, g),
		},
	}, nil
}

func FindPath(start Position, end Position, g [][]int) [][]int {

	numRows := len(g)
	numCols := len(g[0])

	startNode := Node{nil, start, 0, 0, 0}
	endNode := Node{nil, end, 0, 0, 0}

	openList := []Node{startNode}
	closedList := []Node{}

	for len(openList) > 0 {

		currentNode := openList[0]
		currentIndex := 0

		for index, item := range openList {
			if item.f < currentNode.f {
				currentNode = item
				currentIndex = index
			}
		}

		openList = remove(openList, currentIndex)
		closedList = append(closedList, currentNode)

		if currentNode.position.X == endNode.position.X && currentNode.position.Y == endNode.position.Y {
			path := [][]int{}
			current := &currentNode

			for current != nil {
				step := []int{current.position.X, current.position.Y}
				path = append([][]int{step}, path...)
				current = current.parent
			}

			return path
		}

		children := []Node{}
		options := []Position{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

		for _, newPosition := range options {
			nodePosition := addPositions(currentNode.position, newPosition)

			if nodePosition.X >= numRows || nodePosition.X < 0 || nodePosition.Y >= numCols || nodePosition.Y < 0 {
				continue
			}

			if g[nodePosition.X][nodePosition.Y] != 0 {
				continue
			}

			newNode := Node{&currentNode, nodePosition, 0, 0, 0}

			children = append(children, newNode)
		}

		for _, child := range children {

			skip := false
			for _, closedChild := range closedList {
				if closedChild.position.X == child.position.X && closedChild.position.Y == child.position.Y {
					skip = true
					break
				}
			}

			if skip {
				continue
			}

			child.g = currentNode.g + 1
			child.h = calculateH(child.position, endNode.position)
			child.f = child.g + child.h

			for _, openNode := range openList {
				if openNode.position.X == child.position.X && openNode.position.Y == child.position.Y && child.g > openNode.g {
					skip = true
					break
				}
			}

			if skip {
				continue
			}

			openList = append(openList, child)
		}
	}
	return nil
}

func remove(slice []Node, s int) []Node {
	return append(slice[:s], slice[s+1:]...)
}

func addPositions(pos1 Position, pos2 Position) Position {
	return Position{pos1.X + pos2.X, pos1.Y + pos2.Y}
}

func calculateH(pos1 Position, pos2 Position) int {
	x := pos1.X - pos2.X
	y := pos1.Y - pos2.Y
	return (x * x) + (y * y)
}
