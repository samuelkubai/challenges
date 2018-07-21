package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"fmt"
	"io"
)

type Node struct {
	value []int32
	next  *Node
}

type LinkedList struct {
	end    *Node
	start  *Node
	length int
}

func (l *LinkedList) attach(value []int32) {
	node := &Node{value: value}

	oldStart := l.start
	l.start = node
	l.start.next = oldStart
	l.length++
}

func (l *LinkedList) append(value []int32) {
	node := &Node{value: value}

	if l.length < 1 {
		l.start = node
		l.end = node
	} else {
		l.end.next = node
		l.end = node
	}

	l.length++
}

func (l *LinkedList) remove() (removedNode *Node) {
	removedNode = l.start

	if l.start.next != nil {
		l.start = l.start.next
	}

	l.length--

	return removedNode
}

func (l *LinkedList) reset() {
	l.length = 0
}

type Grid struct {
	cells   [][]int32
	visited [][]int32
}

func (g *Grid) first() (int32, int32) {
	return 0, 0
}

func (g *Grid) horizontalLength() int32 {
	return int32(len(g.cells))
}

func (g *Grid) verticalLength() int32 {
	return int32(len(g.cells[0]))
}

func (g *Grid) adjacent(x, y int32) (adjacent [][]int32) {
	var hasRightAdjacent, hasLeftAdjacent, hasTopAdjacent, hasBottomAdjacent = false, false, false, false

	if horizontalAdjacent := x + 1; horizontalAdjacent < g.horizontalLength() {
		adjacent = append(adjacent, []int32{horizontalAdjacent, y})
		hasRightAdjacent = true
	}

	if horizontalAdjacent := x - 1; horizontalAdjacent > -1 {
		adjacent = append(adjacent, []int32{horizontalAdjacent, y})
		hasLeftAdjacent = true
	}

	if verticalAdjacent := y + 1; verticalAdjacent < g.verticalLength() {
		adjacent = append(adjacent, []int32{x, verticalAdjacent})
		hasBottomAdjacent = true
	}

	if verticalAdjacent := y - 1; verticalAdjacent > -1 {
		adjacent = append(adjacent, []int32{x, verticalAdjacent})
		hasTopAdjacent = true
	}


	if hasRightAdjacent && hasBottomAdjacent {
		adjacent = append(adjacent, []int32{x + 1, y + 1})
	}

	if hasLeftAdjacent && hasTopAdjacent {
		adjacent = append(adjacent, []int32{x - 1, y - 1})
	}

	if hasRightAdjacent && hasTopAdjacent {
		adjacent = append(adjacent, []int32{x + 1, y - 1})
	}

	if hasLeftAdjacent && hasBottomAdjacent {
		adjacent = append(adjacent, []int32{x - 1, y + 1})
	}

	return adjacent
}

func (g *Grid) markAsVisited(x, y int32) {
	g.visited = append(g.visited, []int32{x, y})
}

func (g *Grid) notVisited(cell []int32) bool {
	for _, visitedCell := range g.visited {
		if visitedCell[0] == cell[0] && visitedCell[1] == cell[1] {
			return false
		}
	}

	return true
}

func (g *Grid) getAllRegionSizes() (regions []int32) {

	return regions
}

//-----------------------------------------------------------------------------------------
func getAllRegionSizes(g *Grid) (sizes []int32) {
	createdList := &LinkedList{}
	currentRegion := &LinkedList{}

	x, y := g.first()
	createdList.append([]int32{x, y})

	fmt.Printf("Linked list length: %v\n", createdList.length)

	for createdList.length > 0 {
		node := createdList.remove()

		if g.notVisited(node.value) {
			if g.cells[node.value[0]][node.value[1]] == 1 {
				currentRegion.append(node.value)
				fmt.Printf("Filled cell -> %v\n", node.value)
				fmt.Printf("Fill cell -> Adjacent cells %v\n", g.adjacent(node.value[0], node.value[1]))
				for _, location := range g.adjacent(node.value[0], node.value[1]) {
					if g.cells[location[0]][location[1]] == 1 {
						fmt.Printf("Put in the front of the list %v\n", location)
						createdList.attach(location)
					} else {
						createdList.append(location)
					}
				}
			} else {
				fmt.Printf("Empty cell -> %v\n", node.value)
				if currentRegion.length > 0 {
					sizes = append(sizes, int32(currentRegion.length))
					currentRegion.reset()
				}

				adjacentLocations := g.adjacent(node.value[0], node.value[1])
				//fmt.Printf("Empty cell -> Adjacent cells %v\n", adjacentLocations)

				for _, location := range adjacentLocations {
					createdList.append(location)
				}
			}

			g.markAsVisited(node.value[0], node.value[1])
		}
	}

	if currentRegion.length > 0 {
		sizes = append(sizes, int32(currentRegion.length))
		currentRegion.reset()
	}

	return sizes
}
//-----------------------------------------------------------------------------------------

// Complete the maxRegion function below.
func maxRegion(grid [][]int32) (largestRegionSize int32) {
	var regions []int32

	regions = getAllRegionSizes(&Grid{ cells: grid })

	fmt.Printf("Regions found: %v \n", regions)

	for _, regionSize := range regions {
		if regionSize > largestRegionSize {
			largestRegionSize = regionSize
		}
	}

	fmt.Printf("Largest region found: %v \n", largestRegionSize)

	return largestRegionSize
}

type CustomWriter int

func (*CustomWriter) Write(p []byte) (n int, err error) {
	fmt.Printf("Largest region is: %q", p)
	return len(p), nil
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	//stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	//checkError(err)
	//
	//defer stdout.Close()

	//writer := bufio.NewWriterSize(stdout, 1024*1024)
	w := new(CustomWriter)
	writer := bufio.NewWriterSize(w, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	mTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	m := int32(mTemp)

	var grid [][]int32
	for i := 0; i < int(n); i++ {
		gridRowTemp := strings.Split(readLine(reader), " ")

		var gridRow []int32
		for _, gridRowItem := range gridRowTemp {
			gridItemTemp, err := strconv.ParseInt(gridRowItem, 10, 64)
			checkError(err)
			gridItem := int32(gridItemTemp)
			gridRow = append(gridRow, gridItem)
		}

		if len(gridRow) != int(m) {
			panic("Bad input")
		}

		grid = append(grid, gridRow)
	}

	res := maxRegion(grid)

	fmt.Fprintf(writer, "%d\n", res)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
