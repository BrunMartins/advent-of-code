// Day1
package main

import (
	"advent-of-code/common"
	"bufio"
	"os"
)

var (
	puzzleInput                *os.File
	floor                      = 0
	foundBasement              = false
	firstBasementVisitPosition int
)

func getInputLineScanner() *bufio.Scanner {
	fileScanner := bufio.NewScanner(puzzleInput)
	fileScanner.Split(bufio.ScanLines)

	return fileScanner
}

func calculateFloorChanges(input string) {
	for i, char := range input {
		if string(char) == string('(') {
			floor++
		} else if string(char) == string(')') {
			floor--
		}

		if !foundBasement && floor == -1 {
			foundBasement = true
			firstBasementVisitPosition = i + 1
		}
	}
}

func iterateInputFile() {
	scanner := getInputLineScanner()
	for scanner.Scan() {
		calculateFloorChanges(scanner.Text())
	}
}

func main() {
	var err error
	puzzleInput, err = common.OpenPuzzleInput()

	if err != nil {
		panic(err)
	}

	iterateInputFile()

	println("Floor: ", floor)
	println("Fist basement visit: ", firstBasementVisitPosition)
}
