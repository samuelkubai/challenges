// Description Link: https://www.hackerrank.com/challenges/ctci-ice-cream-parlor/problem

package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
	"fmt"
	"sort"
)

type flavor struct {
	cost int32
	id int32
}

type store struct{
	flavors []*flavor
}

func (s *store) addFlavor(cost int32) {
	s.flavors = append(s.flavors, &flavor{ id: int32(len(s.flavors) + 1), cost: cost})
}

func (s *store) getFlavor(index int32) *flavor {
	return s.flavors[index]
}

// Implement the methods below to enable the store
// struct to make use of the in-built sorting
// libraries in go
func (s *store) Len() int {
	return len(s.flavors)
}
func (s *store) Swap(i, j int) {
	s.flavors[i], s.flavors[j] = s.flavors[j], s.flavors[i]
}
func (s store) Less(i, j int) bool {
	return s.flavors[i].cost < s.flavors[j].cost
}
// End of the sorting enabling methods.


// Complete the whatFlavors function below.
func whatFlavors(currentStore *store, money int32) {
	// Sort
	sort.Stable(currentStore)

	// Find pair
	var lowerFlavor *flavor
	var lowerIndex, possibleComplement int32 = 0, 0

	for {
		// If we don't find a match an thus
		// is the end of the loop
		if lowerIndex >= int32(currentStore.Len()) {
			fmt.Printf("No cost pair found")
			return
		}

		lowerFlavor = currentStore.getFlavor(lowerIndex)

		possibleComplement = money - lowerFlavor.cost

		if matchedFlavor, valueFound := findCost(currentStore, possibleComplement, lowerIndex + 1); valueFound {
			// Print the results
			if lowerFlavor.id > matchedFlavor.id {
				fmt.Printf("%d %d\n", matchedFlavor.id, lowerFlavor.id)
			} else {
				fmt.Printf("%d %d\n", lowerFlavor.id, matchedFlavor.id)
			}
			return
		}

		lowerIndex++
	}
}

func findCost(currentStore *store, targetCost int32, startPoint int32) (*flavor, bool) {
	var searchIndex int32

	max := int32(currentStore.Len() - 1)
	min := startPoint

	for {
		if min > max {
			break
		}

		searchIndex = (min + max) / 2
		selectedFlavor := currentStore.getFlavor(searchIndex)

		if selectedFlavor.cost == targetCost {
			return selectedFlavor, true
		}

		if selectedFlavor.cost > targetCost {
			max = searchIndex - 1
		} else {
			min = searchIndex + 1
		}
	}

	return &flavor{}, false
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

		var currentStore = &store{ flavors: make([]*flavor, 0) }

		for i := 0; i < int(n); i++ {
			costItemTemp, err := strconv.ParseInt(costTemp[i], 10, 64)
			checkError(err)
			costItem := int32(costItemTemp)
			currentStore.addFlavor(costItem)
		}

		whatFlavors(currentStore, money)
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
