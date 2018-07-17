package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
	"fmt"
)

// Complete the whatFlavors function below.
func whatFlavors(cost []int32, money int32) {
	// Sort the costs in order to allow us to
	// be able to search for the costs
	sortedCosts, originalIndicies := mergeSort(cost, 0)
	//fmt.Printf("Sorted costs are %v\n", sortedCosts)
	//fmt.Printf("Sorted indicies are %v\n", originalIndicies)

	// Chunk the costs slice in order to get
	// a small slice to work on.
	minimizedCosts := chunkCost(sortedCosts, money)
	//fmt.Printf("Chunked cost slice is %v\n", minimizedCosts)

	var lowerIndex, lowerCost, estimatedHigherCost int32 = 0, 0, 0

	for {
		// If we don't find a match an thus
		// is the end of the loop
		if lowerIndex >= int32(len(minimizedCosts)) {
			break
		}

		lowerCost = minimizedCosts[lowerIndex]
		estimatedHigherCost = money - lowerCost

		if higherIndex, valueFound := findCost(minimizedCosts, estimatedHigherCost, 0); valueFound {
			resultLow := originalIndicies[lowerIndex] + 1
			resultHigh := originalIndicies[higherIndex] + 1

			if resultLow > resultHigh {
				resultHigh = originalIndicies[lowerIndex] + 1
				resultLow = originalIndicies[higherIndex] + 1
			}

			fmt.Printf("%d %d\n", resultLow, resultHigh)
			return
		}

		lowerIndex++
	}

	fmt.Printf("No cost pair found")
	return
}

func findCost(costs []int32, targetCost int32, startIndex int32) (int32, bool) {
	if len(costs) < 2 {
		if costs[0] == targetCost {
			return startIndex, true
		}

		return -1, false
	}

	middle := int32(len(costs) / 2)

	if costs[middle] == targetCost {
		return startIndex + middle, true
	}

	if len(costs) < 2 {
		return -1, false
	}

	if costs[middle] > targetCost {
		return findCost(costs[:middle], targetCost, startIndex)
	} else {
		return findCost(costs[middle:], targetCost, middle)
	}
}

func chunkCost(costs []int32, money int32) []int32 {
	middle := int32(len(costs) / 2)

	// Check if the money is initially larger
	// than any single cost provided and
	// just return the whole cost
	// slice since any of the
	// cost is a viable
	// candidate.
	if money >= costs[len(costs) - 1] {
		return costs
	}

	if middle == money {
		return costs[:money]
	}

	if middle > money {
		return chunkCost(costs[:middle], money)
	}

	return costs
}

func mergeSort(slice []int32, originalIndex int32) ([]int32, []int32) {
	if len(slice) < 2 {
		return slice, []int32{originalIndex}
	}

	middle := int32(len(slice) / 2)
	leftItems, leftIndicies := mergeSort(slice[:middle], originalIndex)
	rightItems, rightIndicies := mergeSort(slice[middle:], originalIndex + middle)
	return merge(leftItems, rightItems, leftIndicies, rightIndicies)
}

func merge(left []int32, right []int32, leftIndicies []int32, rightIndicies []int32) ([]int32, []int32)  {
	var leftIndex, rightIndex int32 = 0, 0
	var temp []int32
	var tempIndicies []int32
	for {
		if leftIndex >= int32(len(left)) {
			temp = append(temp, right[rightIndex:]...)
			tempIndicies = append(tempIndicies, rightIndicies[rightIndex:]...)
			break
		}

		if rightIndex >= int32(len(right)) {
			temp = append(temp, left[leftIndex:]...)
			tempIndicies = append(tempIndicies, leftIndicies[leftIndex:]...)
			break
		}

		if left[leftIndex] <= right[rightIndex] {
			temp = append(temp, left[leftIndex])
			tempIndicies = append(tempIndicies, leftIndicies[leftIndex])
			leftIndex++
		} else {
			temp = append(temp, right[rightIndex])
			tempIndicies = append(tempIndicies, rightIndicies[rightIndex])
			rightIndex++
		}
	}

	return temp, tempIndicies
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		moneyTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		money := int32(moneyTemp)

		nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		n := int32(nTemp)

		costTemp := strings.Split(readLine(reader), " ")

		var cost []int32

		for i := 0; i < int(n); i++ {
			costItemTemp, err := strconv.ParseInt(costTemp[i], 10, 64)
			checkError(err)
			costItem := int32(costItemTemp)
			cost = append(cost, costItem)
		}

		whatFlavors(cost, money)
	}
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
