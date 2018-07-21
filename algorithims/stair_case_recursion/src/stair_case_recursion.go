// Challenge: https://www.hackerrank.com/challenges/ctci-recursive-staircase/problem
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type StepsPremsCollection struct{
	memory map[int32]int32
}

func (s *StepsPremsCollection) placeInMemory(value, result int32) {
	s.memory[value] = result
}

func (s *StepsPremsCollection) remember(value int32) (bool, int32) {
	if memoryValue := s.memory[value]; memoryValue > 0 {
		return true, memoryValue
	}

	return false, 0
}

func (s *StepsPremsCollection) getPerms(value int32) int32 {
	// Base line
	if value <= 0 {
		return 0
	}
	if value == 1 {
		return 1
	}
	if value == 2 {
		return 2
	}
	if value == 3 {
		return 4
	}

	// Check if the result for the value provided
	// has already been computed and retrieve
	// and avoid extra computation.
	if exists, result := s.remember(value); exists {
		return result
	}

	result := s.getPerms(value-1) + s.getPerms(value-2) + s.getPerms(value-3)
	s.placeInMemory(value, result)

	return result
}

// Complete the stepPerms function below.
func stepPerms(n int32) int32 {
	collection := &StepsPremsCollection{memory: make(map[int32]int32)}
	return collection.getPerms(n)
}


type CustomWriter int

func (*CustomWriter) Write(p []byte) (n int, err error) {
	fmt.Printf("%s", p)
	return len(p), nil
}


func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	//stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	//checkError(err)
	//
	//defer stdout.Close()

	//writer := bufio.NewWriterSize(stdout, 1024*1024)
	w := new(CustomWriter)
	writer := bufio.NewWriterSize(w, 1024 * 1024)

	sTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	s := int32(sTemp)

	for sItr := 0; sItr < int(s); sItr++ {
		nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		n := int32(nTemp)

		res := stepPerms(n)

		fmt.Fprintf(writer, "%d\n", res)
	}

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
