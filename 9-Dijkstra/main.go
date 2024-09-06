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
	effort int
	x      int
	y      int
}

type PriorityQueue []Item

func (p PriorityQueue) Len() int { return len(p) }

func (p PriorityQueue) Less(i, j int) bool {
	return p[i].effort < p[j].effort
}

func (p PriorityQueue) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *PriorityQueue) Push(x interface{}) {
	*p = append(*p, x.(Item))
}

func (p *PriorityQueue) Pop() interface{} {
	old := *p
	n := len(old)
	item := old[n-1]
	*p = old[0 : n-1]
	return item
}

func minimumEffortPath(heights [][]int) int {
	rows, columns := len(heights), len(heights[0])
	distance := make([][]int, rows)
	for i := range distance {
		distance[i] = make([]int, columns)
		for j := range distance[i] {
			distance[i][j] = math.MaxInt32
		}
	}
	distance[0][0] = 0
	directions := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	p := &PriorityQueue{Item{0, 0, 0}}
	heap.Init(p)

	for p.Len() > 0 {
		item := heap.Pop(p).(Item)
		effort, x, y := item.effort, item.x, item.y
		if effort > distance[x][y] {
			continue
		}
		if x == rows-1 && y == columns-1 {
			return effort
		}
		for _, dir := range directions {
			nx, ny := x+dir[0], y+dir[1]
			if nx >= 0 && nx < rows && ny >= 0 && ny < columns {
				newEffort := int(math.Max(float64(effort), math.Abs(float64(heights[x][y]-heights[nx][ny]))))
				if newEffort < distance[nx][ny] {
					distance[nx][ny] = newEffort
					heap.Push(p, Item{newEffort, nx, ny})
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
	columns, _ := strconv.Atoi(strings.TrimSpace(readLine(reader)))

	heights := make([][]int, rows)
	for i := 0; i < rows; i++ {
		fmt.Printf("enter row %d: ", i+1)
		line := strings.TrimSpace(readLine(reader))
		values := strings.Split(line, " ")
		heights[i] = make([]int, columns)
		for j := 0; j < columns; j++ {
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
