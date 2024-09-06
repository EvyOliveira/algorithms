package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Item struct {
	effort, x, y int
}

type PriorityQueue []Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].effort < pq[j].effort
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(Item))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func minimumEffortPath(heights [][]int) int {
	rows, cols := len(heights), len(heights[0])
	dist := make([][]int, rows)
	for i := range dist {
		dist[i] = make([]int, cols)
		for j := range dist[i] {
			dist[i][j] = math.MaxInt32
		}
	}
	dist[0][0] = 0
	directions := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	pq := &PriorityQueue{Item{0, 0, 0}}
	heap.Init(pq)

	for pq.Len() > 0 {
		item := heap.Pop(pq).(Item)
		effort, x, y := item.effort, item.x, item.y
		if effort > dist[x][y] {
			continue
		}
		if x == rows-1 && y == cols-1 {
			return effort
		}
		for _, dir := range directions {
			nx, ny := x+dir[0], y+dir[1]
			if nx >= 0 && nx < rows && ny >= 0 && ny < cols {
				newEffort := int(math.Max(float64(effort), math.Abs(float64(heights[x][y]-heights[nx][ny]))))
				if newEffort < dist[nx][ny] {
					dist[nx][ny] = newEffort
					heap.Push(pq, Item{newEffort, nx, ny})
				}
			}
		}
	}
	return -1
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("enter the number of rows: ")
	rows, _ := strconv.Atoi(strings.TrimSpace(readLine(reader)))

	fmt.Print("enter the number of columns: ")
	cols, _ := strconv.Atoi(strings.TrimSpace(readLine(reader)))

	heights := make([][]int, rows)
	for i := 0; i < rows; i++ {
		fmt.Printf("enter row %d: ", i+1)
		line := strings.TrimSpace(readLine(reader))
		values := strings.Split(line, " ")
		heights[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			heights[i][j], _ = strconv.Atoi(values[j])
		}
	}

	minEffort := minimumEffortPath(heights)
	fmt.Printf("minimum effort: %d\n", minEffort)
}

func readLine(reader *bufio.Reader) string {
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}
